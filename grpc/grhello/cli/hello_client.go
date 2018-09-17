package main

import (
	"google.golang.org/grpc"
	"log"
	pb "gin-lnp-api/grpc/grhello"
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
	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Greet(ctx,
		&pb.Request{FirstName: "Ryan",
			LastName: "Tran",})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", response.Greeting)
}


