package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tclohm/grpc-playground/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from the client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	//fmt.Printf("Created client: %f", c)

	unaryRequest(c)

}

func unaryRequest(c greetpb.GreetServiceClient) {
	fmt.Println("starting unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Taylor",
			LastName: "Lohman",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}