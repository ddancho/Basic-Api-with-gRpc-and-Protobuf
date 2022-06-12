package main

import (
	"context"
	"fmt"
	"tensor-grpc-tut/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewAddServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	add, err := client.Add(ctx, &proto.Request{A: 10, B: 10})
	if err != nil {
		panic(err)
	}

	fmt.Println(add.Result)
}
