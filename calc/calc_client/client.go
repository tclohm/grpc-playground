package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tclohm/grpc-playground/calc/calcpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from the client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := calcpb.NewGreetServiceClient(conn)

	//fmt.Printf("Created client: %f", c)

	unaryRequest(c)

}

func unaryRequest(c calcpb.SumServiceClient) {
	fmt.Println("starting unary RPC...")
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			firstNumber: 10,
			SecondNumber: 2,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}