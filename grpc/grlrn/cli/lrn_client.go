package main

import (
	"google.golang.org/grpc"
	"log"
	pb "gin-lnp-api/grpc/grlrn"
	"time"
	"context"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewLrnServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.FindByDid(ctx,
		&pb.DidReq{Did: "7143174661"})

	if err != nil {
		log.Fatalf("could not find: %v", err)
	}

	log.Printf("LRN: %s", response)
}


