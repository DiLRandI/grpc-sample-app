package main

import (
	"context"
	"fmt"
	"net"

	"github.com/dilrandi/grpc-sample-app/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, r *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("\n\n Greet function was invoked %v \n\n", r)
	firstName := r.GetGreeting().GetFirstName()
	lastName := r.GetGreeting().GetLastName()
	result := fmt.Sprintf("Hello Mr. %s %s", firstName, lastName)
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server is running !")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
