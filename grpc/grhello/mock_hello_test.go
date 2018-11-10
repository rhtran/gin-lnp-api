package grhello

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"testing"
	"github.com/golang/mock/gomock"
	"time"
	"context"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestGreet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := NewMockHelloServiceClient(ctrl)
	req := &Request{FirstName: "Ryan", LastName: "Tran"}
	mockGreeterClient.EXPECT().Greet(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&Response{Greeting: "Hello Ryan Tran"}, nil)
	testGreet(t, mockGreeterClient)
}

func testGreet(t *testing.T, client HelloServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Greet(ctx, &Request{FirstName: "Ryan", LastName: "Tran"})
	if err != nil || r.Greeting != "Hello Ryan Tran" {
		t.Errorf("greeting failed")
	}
	t.Log("Reply : ", r.Greeting)
}