package main

import (
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

	fmt.Printf("Created client: %f", c)
}