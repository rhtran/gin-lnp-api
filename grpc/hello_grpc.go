package grpc

import "golang.org/x/net/context"

type GrpcHelloService struct {

}

func (s *GrpcHelloService) Greet(ctx context.Context, request *Request) (*Response, error) {
	greeting := "Hello " + request.GetFirstName() + " " + request.GetLastName()
	return &Response{Greeting:greeting, Sentiment: request.GetSentiment()}, nil
}
