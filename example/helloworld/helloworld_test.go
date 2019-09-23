package helloworld

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestNewGreeterClient(t *testing.T) {
	conn, err := grpc.Dial(":8032", grpc.WithInsecure())
	if err != nil {
		panic(err)
		return
	}
	res, err := NewGreeterClient(conn).SayHello(context.Background(), &HelloRequest{
		Name: "jack",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

}
