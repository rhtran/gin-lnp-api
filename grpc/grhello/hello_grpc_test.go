package grhello

import (
	"testing"
	"context"
)

func HelloTest(t *testing.T) {
	s := GrpcHelloService{}

	// set up test cases
	tests := []struct{
		firstName string
		lastName string
		want string
	} {
		{
			firstName: "world",
			lastName: "!",
			want: "Hello world !",
		},
		{
			firstName: "123",
			lastName: "456",
			want: "Hello 123 456",
		},
	}

	for _, tt := range tests {
		req := Request{FirstName: tt.firstName,
						LastName: tt.lastName,
		}
		resp, err := s.Greet(context.Background(), &req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error")
		}
		if resp.Greeting != tt.want {
			t.Errorf("HelloText(%v)=%v %v, wanted %v", tt.firstName, tt.lastName, resp.Greeting, tt.want)
		}
	}
}

