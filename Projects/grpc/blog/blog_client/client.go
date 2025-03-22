package main

import (
	"context"
	"fmt"
	"github.com/afasari/go-workspace/GRPC/blog/blogpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Afasari",
		Title:    "Hello World!",
		Content:  "Welcome to the world",
	}
	createBlogRes, createBlogErr := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog,})
	if createBlogErr != nil {
		log.Fatalf("Unexpected error : %v", createBlogErr)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// read Blog
	fmt.Println("Reading the blog")
	_, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "idfsdfasdf",
	})
	if readBlogErr != nil {
		log.Printf("Error happened while reading: %v\n", readBlogErr)
	}

	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogID,
	})
	if readBlogErr != nil {
		log.Printf("Error happened while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog
	fmt.Println("Updating the blog")
	newBlog := &blogpb.Blog{
		Id:                   blogID,
		AuthorId:             "Change Author",
		Title:                "My First Blog",
		Content:              "Content of first blog",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog:newBlog,
	})
	if updateErr != nil {
		log.Printf("Error happened while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog was update: %v\n", updateRes)

	// delete Blog
	fmt.Println("Deleting the blog")
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId:blogID,
	})
	if deleteErr != nil {
		log.Printf("Error happened while deleting: %v\n", deleteErr)
	}
	fmt.Printf("Blog was delete: %v\n", deleteRes)

	// list blog
	fmt.Println("Listing the blog")
	stream, listErr := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if listErr != nil {
		log.Printf("Error happened while listBlog RPC: %v\n", listErr)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while stream: %v\n", err)
		}
		fmt.Println(res.GetBlog())
	}
}
