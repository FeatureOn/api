package grpc

import (
	"context"
	"log"

	pb "github.com/FeatureOn/api/flagpb"
	"github.com/FeatureOn/api/server/adapters/comm/grpc/mappers"
	"github.com/FeatureOn/api/server/application"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedFlagServiceServer
	flagRepo    application.FlagRepository
	productRepo application.ProductRepository
}

// NewServer creates a new grpc.Server and initializes it with repositories
func NewServer(flagRepository application.FlagRepository, productRepository application.ProductRepository) *Server {
	return &Server{
		flagRepo:    flagRepository,
		productRepo: productRepository,
	}
}

// GetEnvironmentFlags implements featureon.api.flagpb.GetEnvironmentFlags
func (s *Server) GetEnvironmentFlags(ctx context.Context, in *pb.EnvironmentFlagQuery) (*pb.EnvironmentFlags, error) {
	log.Printf("Received: %v", in.GetEnvironmentID())
	productService := application.NewProductService(s.productRepo, s.flagRepo)
	envFlag, err := productService.GetFlags(in.GetEnvironmentID())
	if err != nil {
		return &pb.EnvironmentFlags{}, err
	}
	return mappers.MapEnvironmentFlag2EnvironmentFlags(envFlag), nil
}
