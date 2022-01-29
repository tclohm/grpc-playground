package main

import (
	"fmt"
	"net"
	"log"

	"github.com/tclohm/grpc-playground/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct {}

func main() {
	fmt.Println("Hello world!")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}