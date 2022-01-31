package main

import (
	"context"
	"fmt"
	"log"
	"io"

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

	c := calcpb.NewSumServiceClient(conn)

	//fmt.Printf("Created client: %f", c)

	//unaryRequest(c)

	serverStreaming(c)

}

func unaryRequest(c calcpb.SumServiceClient) {
	fmt.Println("starting unary RPC...")
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			FirstNumber: 10,
			SecondNumber: 2,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Result)
}

func serverStreaming(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &calcpb.SumManyTimesRequest{
		Sum: &calcpb.Sum{
			FirstNumber: 120,
		},
	}

	res, err := c.SumManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling SumManyTimes RPC: %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			fmt.Println("End")
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from SumManyTimes: %v", msg.GetResult())
	}
}