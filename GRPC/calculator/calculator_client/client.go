package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/afasari/go-workspace/GRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a Client")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	// doBidiStreaming(c)

	doErrorUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 40,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v\n", err)
	}
	log.Printf("Response from Calculator: %v\n", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a server streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculator RPC: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while stream: %v\n", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a client streaming RPC...")
	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while opening stream %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v\n", err)
	}
	fmt.Printf("The average is: %v\n", res.GetAverage())

}

func doBidiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a bidi streaming RPC...")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while opening stream and calling FindMaximum: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		numbers := []int32{3, 2, 5, 10, 24, 2, 48}
		for _, number := range numbers {
			fmt.Printf("Sending number: %v\n", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Problem while reading server stream: %v\n", err)
				break
			}
			max := res.GetNumber()
			fmt.Printf("Receive a new max number: %v\n", max)
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot unary RPC...")
	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -10)

}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {

	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Println("Error Message from server:", resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably send a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
}
