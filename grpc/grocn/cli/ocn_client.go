package main

import (
	"google.golang.org/grpc"
	"time"
	"log"
	pb "gin-lnp-api/grpc/grocn"
	"context"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOcnServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.FindByOcn(ctx,
		&pb.OcnReq{Ocn: "1583"})

	if err != nil {
		log.Fatalf("could not get OCN: %v", err)
	}

	log.Printf("Greeting: %s", response)
}

