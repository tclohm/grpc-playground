package main

import (
	"context"
	"fmt"
	"net"
	"log"
	//"time"

	"github.com/tclohm/grpc-playground/calc/calcpb"

	"google.golang.org/grpc"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	log.Printf("Sum invoked with %v", req)

	firstNumber := req.GetSum().GetFirstNumber()
	SecondNumber := req.GetSum().GetSecondNumber()

	result := firstNumber + SecondNumber

	res := &calcpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func (*server) SumManyTimes(req *calcpb.SumManyTimesRequest, stream calcpb.SumService_SumManyTimesServer) error {
	log.Printf("Sum many times invoked with %v", req)

	var firstNumber int32 = req.GetSum().GetFirstNumber()
	
	var constant int32 = 2

	for firstNumber > 1  {
		if firstNumber % constant == 0 {
			
			res := &calcpb.SumManyTimesResponse{
				Result: constant,
			}
			stream.Send(res)
			firstNumber = firstNumber / constant
		} else {
			constant += 1
		}
	}

	return nil
}

func main() {
	fmt.Println("Server on")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}