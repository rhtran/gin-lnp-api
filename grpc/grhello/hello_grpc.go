package grhello

//go:generate mockgen -destination=grhello/mock_hello -source=grhello/hello_grpc.go -package=grhello

import (
	"golang.org/x/net/context"
	)

type GrpcHelloService struct {
}

func NewGrpcHelloService() *GrpcHelloService {
	return &GrpcHelloService {
	}
}

func (s *GrpcHelloService) Greet(ctx context.Context, request *Request) (*Response, error) {
	greeting := "Hello " + request.GetFirstName() + " " + request.GetLastName()
	return &Response{Greeting:greeting, Sentiment: request.GetSentiment()}, nil
}
