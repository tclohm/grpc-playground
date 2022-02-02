package main

import (
	"log"
	"net"
	"fmt"
	"os"
	"os/signal"
	"context"

	

	"github.com/tclohm/grpc-playground/blog/blogpb"

	"google.golang.org/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID 			int32 	`bson:"_id,omitempty"`
	AuthorID 	string	`bson:"author_id"`
	Content 	string 	`bson:"content"`
	Title 		string 	`bson:"title"`
}

func main() {

	// if we crash, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog Service Started")
	fmt.Println("Connecting to mongodb")
	// connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { log.Fatal(err) }

	err = client.Connect(context.TODO())
	if err != nil { log.Fatal(err) }

	fmt.Println("Connecting to collection")
	collection = client.Database("myweblog").Collection("blog")


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
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}