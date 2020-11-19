package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	rest "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest"
	//memory "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/memory"
	mongodb "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/mongodb"
	"github.com/nicholasjackson/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	util "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/util"
)

var bindAddress = env.String("BASE_URL", false, ":5500", "Bind address for the server")

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	util.SetConstValues()
	util.SetLogLevels()

	//dbContext := memory.NewDataContext()
	dbContext := mongodb.NewDataContext()
	//s := rest.NewAPIContext(dbContext, bindAddress)
	s := rest.NewAPIContext(bindAddress, dbContext.HealthRepository, dbContext.UserRepository)

	// start the server
	go func() {
		log.Debug().Msgf("Starting server on %s", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			log.Error().Err(err).Msg("Error starting server")
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
