package main

import (
	"context"
	"log"
	"time"

	pb "github.com/FeatureOn/api/flagpb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5501"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(20)*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFlagServiceClient(conn)

	// Contact the server and print out its response.
	// name := defaultName
	// if len(os.Args) > 1 {
	// 	name = os.Args[1]
	// }
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20)*time.Second)
	defer cancel()
	r, err := c.GetEnvironmentFlags(ctx, &pb.EnvironmentFlagQuery{EnvironmentID: "5ffcae0fd055b0d1ea6de4f4"})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Result: %v", r)
}
