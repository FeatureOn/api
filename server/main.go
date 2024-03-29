package main

import (
	"context"
	"github.com/FeatureOn/api/server/adapters/data"
	"github.com/FeatureOn/api/server/adapters/data/cockroachdb"
	"github.com/FeatureOn/api/server/adapters/data/mongodb"
	"net"
	"os"
	"os/signal"
	"time"

	pb "github.com/FeatureOn/api/flagpb"
	grpcServer "github.com/FeatureOn/api/server/adapters/comm/grpc"
	"github.com/FeatureOn/api/server/adapters/comm/rest"
	"google.golang.org/grpc"

	"github.com/nicholasjackson/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/FeatureOn/api/server/util"
)

var bindAddress = env.String("BASE_URL", false, ":5500", "Bind address for rest server")
var grpcAddress = env.String("GRPC_URL", false, ":5501", "Bind address for grpc server")
var dbType = env.String("DB_TYPE", false, "cockroachdb", "The preferred database provider. possible values: mongodb, cockroachdb. default: cockroachdb")

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	util.SetConstValues()
	util.SetLogLevels()

	var dbContext data.DataContext
	switch *dbType {
	case "mongodb":
		dbContext = mongodb.NewDataContext()
	default:
		dbContext = cockroachdb.NewDataContext()
	}
	s := rest.NewAPIContext(bindAddress, dbContext.HealthRepository, dbContext.UserRepository, dbContext.ProductRepository, dbContext.FlagRepository)

	// start the http server
	go func() {
		log.Debug().Msgf("Starting server on %s", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			log.Error().Err(err).Msg("Error starting rest server")
			os.Exit(1)
		}
	}()

	g := grpcServer.NewServer(dbContext.FlagRepository, dbContext.ProductRepository)

	// start the grpc server
	go func() {
		lis, err := net.Listen("tcp", *grpcAddress)
		if err != nil {
			log.Error().Err(err).Msg("Error starting grpc server")
			os.Exit(1)
		}
		s := grpc.NewServer()
		pb.RegisterFlagServiceServer(s, g)
		if err := s.Serve(lis); err != nil {
			log.Error().Err(err).Msg("Error starting grpc server")
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Info().Msgf("Got signal: %s", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
