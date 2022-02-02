package main

import (
	"log"
	"net"
	"fmt"
	"os"
	"os/signal"

	"github.com/tclohm/grpc-playground/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {

	// if we crash, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog Service Started")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	// tls := false

	// if tls {
	// 	certFile := "ssl/server.crt"
	// 	keyFile := "ssl/server.pem"

	// 	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

	// 	if sslErr != nil {
	// 		log.Fatalf("Failed loading certificates: %v", sslErr)
	// 		return
	// 	}

	// 	opts = append(opts, grpc.Creds(creds))
	// }

	s := grpc.NewServer(opts...)

	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block on Channel until signal
	<- ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	listener.Close()

	fmt.Println("End of Program")
}