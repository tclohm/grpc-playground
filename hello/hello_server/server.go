package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"strconv"
	"time"

	"github.com/tclohm/grpc-playground/hello/hellopb"

	"google.golang.org/grpc"
)

type server struct {}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Printf("Hello invoked with %v", req)
	firstName := req.GetHello().GetFirstName()

	result := "Hello " + firstName

	res := &hellopb.HelloResponse{
		Result: result,
	}

	return res, nil
}

func (*server) HelloManyTimes(req *hellopb.HelloManyTimesRequest, stream hellopb.HelloService_HelloManyTimesServer) error {
	fmt.Printf("Greet many times function invoked with %v\n", req)
	firstName := req.GetHello().GetFirstName()
	for i := 0 ; i < 10 ; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &hellopb.HelloManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello world!")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	
	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}