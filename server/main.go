package main

import (
	"context"
	"net"
	"tensor-grpc-tut/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AddServiceServer struct {
	proto.UnimplementedAddServiceServer
}

/*
	go mod init <name of module>
	go mod tidy
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
*/

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &AddServiceServer{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *AddServiceServer) Add(ctx context.Context, requets *proto.Request) (*proto.Response, error) {
	a, b := requets.GetA(), requets.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *AddServiceServer) Multiply(ctx context.Context, requets *proto.Request) (*proto.Response, error) {
	a, b := requets.GetA(), requets.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
