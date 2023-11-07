package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/ablingchos/grpc-learning/pkg/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := "pakin"
	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("greeting: %s", res.Message)
}
