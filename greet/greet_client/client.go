package main

import (
	"context"
	"fmt"

	"github.com/dilrandi/grpc-sample-app/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client!")
	con, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer con.Close()
	c := greetpb.NewGreetServiceClient(con)

	// fmt.Printf("Client created %+f \n", c)
	DoUnary(c)
}

func DoUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Deleema",
			LastName:  "Fernando",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n Response is  %v \n\n", res)
}
