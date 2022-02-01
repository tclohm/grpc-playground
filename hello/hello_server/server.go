package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"strconv"
	"time"
	"io"

	"github.com/tclohm/grpc-playground/hello/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
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

func (*server) LongHello(stream hellopb.HelloService_LongHelloServer) error {

	fmt.Printf("Long hello invoked with streaming request\n")

	var result string = ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finished reading client stream
			return stream.SendAndClose(&hellopb.LongHelloResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		firstName := req.GetHello().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) HelloEveryone(stream hellopb.HelloService_HelloEveryoneServer) error {
	fmt.Printf("Hello everyone invoked with bi-dir streaming")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstname := req.GetHello().GetFirstName()

		result := "Hello " + firstname + "!"

		err = stream.Send(&hellopb.HelloEveryoneResponse{
			Result: result,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}

func (*server) HelloWithDeadline(ctx context.Context, req *hellopb.HelloWithDeadlineRequest) (*hellopb.HelloWithDeadlineResponse, error) {
	log.Printf("Hello invoked with %v", req)
	for i := 0 ; i < 3 ; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("Client canceled request")
			return nil, status.Error(codes.DeadlineExceeded, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetHello().GetFirstName()

	result := "Hello " + firstName

	res := &hellopb.HelloWithDeadlineResponse{
		Result: result,
	}

	return res, nil
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