package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tclohm/grpc-playground/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	fmt.Println("Creating Blog")
	blog := &blogpb.Blog{
		AuthorId: "Taylor",
		Title: "Wow wow we wow",
		Content: "Content matters",
	}

	created, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Error in creating blog: %v", err)
	}
	fmt.Printf("Blog created: %v", created)


	fmt.Println("\nReading the blog")

	read, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: created.GetBlog().GetId()})
	if err != nil {
		fmt.Printf("Error while reading: %v", err)
	}

	fmt.Printf("Blog asked for %v", read)
}