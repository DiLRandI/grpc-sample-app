package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dilrandi/grpc-sample-app/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client starting")

	con, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error occurred : %s\n", err.Error())
	}

	defer con.Close()

	c := calculatorpb.NewCalculatorServiceClient(con)

	r := &calculatorpb.CalcRequest{
		No1: 5,
		No2: 2,
	}
	res, err := c.Sum(context.Background(), r)

	if err != nil {
		log.Fatalf("Error occurred : %s\n", err.Error())
	}

	fmt.Println("=======================================================================")
	fmt.Printf("\n\n %+v \n\n", res)
	fmt.Println("=======================================================================")
}