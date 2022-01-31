package main

import (
	"context"
	"fmt"
	"log"
	"io"
	"time"

	"github.com/tclohm/grpc-playground/hello/hellopb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello from the client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := hellopb.NewHelloServiceClient(conn)

	//fmt.Printf("Created client: %f", c)

	//unaryRequest(c)

	//serverStreaming(c)

	//clientStreaming(c)

	biStreaming(c)

}

func unaryRequest(c hellopb.HelloServiceClient) {
	fmt.Println("starting unary RPC...")
	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Taylor",
			LastName: "Lohman",
		},
	}

	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling hello RPC: %v", err)
	}

	log.Printf("Response from hello: %v", res.Result)
}

func serverStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")
	req := &hellopb.HelloManyTimesRequest{
		Hello: &hellopb.Hello{
			FirstName: "Parker",
			LastName: "Lohman",
		},
	}

	res, err := c.HelloManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling HelloManyTimes RPC: %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			// reached end
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from HelloManyTimes: %v", msg.GetResult())
	}

}

func clientStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")


	requests := []*hellopb.LongHelloRequest{
		&hellopb.LongHelloRequest{
			Hello: &hellopb.Hello{
				FirstName: "Taylor",
			},
		},

		&hellopb.LongHelloRequest{
			Hello: &hellopb.Hello{
				FirstName: "Parker",
			},
		},

		&hellopb.LongHelloRequest{
			Hello: &hellopb.Hello{
				FirstName: "Marta",
			},
		},

		&hellopb.LongHelloRequest{
			Hello: &hellopb.Hello{
				FirstName: "Mark",
			},
		},

		&hellopb.LongHelloRequest{
			Hello: &hellopb.Hello{
				FirstName: "Janet",
			},
		},
	}


	stream, err := c.LongHello(context.Background())

	if err != nil {
		log.Fatalf("error while calling long hello %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending request %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from long hello: %v", err)
	}

	fmt.Printf("Long hello response: %v\n", res)
}

func biStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting bi-directional streaming RPC...")

	// create stream
	stream, err := c.HelloEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: $v", err)
		return
	}

	requests := []*hellopb.HelloEveryoneRequest{
		&hellopb.HelloEveryoneRequest{
			Hello: &hellopb.Hello{
				FirstName: "Taylor",
			},
		},

		&hellopb.HelloEveryoneRequest{
			Hello: &hellopb.Hello{
				FirstName: "Parker",
			},
		},

		&hellopb.HelloEveryoneRequest{
			Hello: &hellopb.Hello{
				FirstName: "Marta",
			},
		},

		&hellopb.HelloEveryoneRequest{
			Hello: &hellopb.Hello{
				FirstName: "Mark",
			},
		},

		&hellopb.HelloEveryoneRequest{
			Hello: &hellopb.Hello{
				FirstName: "Janet",
			},
		},
	}

	waitc := make(chan struct{})

	// send message
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				close(waitc)
			}

			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until done
	<-waitc
}