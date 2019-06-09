package main

import (
	"fmt"
	"net"

	"github.com/dilrandi/grpc-sample-app/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

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
