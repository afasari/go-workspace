package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/afasari/go-workspace/GRPC/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string            `bson:"author_id"`
	Content  string            `bson:"content"`
	Title    string            `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Create blog request")
	blog := req.GetBlog()
	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}
	res ,err := collection.InsertOne(context.Background(), data)
	if err != nil{
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v/n", err),
			)
	}

	objectID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil,status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can't convert to OID : %v/n", err),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:                   objectID.Hex(),
			AuthorId:             blog.GetAuthorId(),
			Title:                blog.GetTitle(),
			Content:              blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error){
	fmt.Println("Read blog request")
	blogID := req.GetBlogId()
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil{
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Can't parse ID"),)
	}

	// create an empty struct
	data := &blogItem{}
	filter := bson.M{"_id": objectID}
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil{
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Can't find blog with the specific ID: %v\n", err))
	}

	return &blogpb.ReadBlogResponse{Blog:dataToBlogPb(data)}, nil
}

func dataToBlogPb(data *blogItem) *blogpb.Blog{
	return &blogpb.Blog{
			Id:                   data.ID.Hex(),
			AuthorId:             data.AuthorID,
			Title:                data.Title,
			Content:              data.Content,
			}
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	fmt.Println("Update blog request")
	blog := req.GetBlog()
	objectID, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil{
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Can't parse ID"),)
	}

	// create an empty struct
	data := &blogItem{}
	filter := bson.M{"_id": objectID}
	res := collection.FindOne(context.Background(), filter)

	if err := res.Decode(data); err != nil{
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Can't find blog with the specific ID: %v\n", err))
	}

	data.AuthorID = blog.GetAuthorId()
	data.Title = blog.GetTitle()
	data.Content = blog.GetContent()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
	if updateErr != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Can't update object in Mongodb: %v\n", updateErr))
	}

	return &blogpb.UpdateBlogResponse{Blog:dataToBlogPb(data),}, nil
}

func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	fmt.Println("Delete blog request")
	blogID := req.GetBlogId()
	objectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil{
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Can't parse ID"),)
	}

	filter := bson.M{"_id": objectID}
	deleteRes, deleteErr := collection.DeleteOne(context.Background(), filter)
	if deleteErr != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Can't delete object in Mongodb: %v\n", deleteErr))
	}

	if deleteRes.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Can't find blog in Mongodb\n"))
	}

	return &blogpb.DeleteBlogResponse{BlogId:blogID}, nil
}

func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	fmt.Println("List blog request")
	filter := bson.D{{}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error find: %v\n", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &blogItem{}
		err = cursor.Decode(data)
		if err != nil{
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data in Mongodb: %v\n", err))
		}
		stream.Send(&blogpb.ListBlogResponse{
			Blog:dataToBlogPb(data),
		})
	}
	if err = cursor.Err(); err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}

	return nil
}

func main() {
	// if we crash the code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Printf("Connecting Mongodb\n")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("mydb").Collection("blog")

	fmt.Printf("Blog Service Started\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	// register reflection service on grpc server
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// wait for control c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing the mongodb connection")
	client.Disconnect(ctx)
	fmt.Println("End of the program")
}
