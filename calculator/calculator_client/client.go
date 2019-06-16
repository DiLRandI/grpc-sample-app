package main

import (
	"context"
	"fmt"
	"io"
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

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	s, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while connecting to server for ComputeAverage : %+v \n", err)
	}

	request := []calculatorpb.ComputeAverageRequest{
		calculatorpb.ComputeAverageRequest{
			Number: 2,
		},
		calculatorpb.ComputeAverageRequest{
			Number: 3,
		},
		calculatorpb.ComputeAverageRequest{
			Number: 5,
		},
		calculatorpb.ComputeAverageRequest{
			Number: 8,
		},
		calculatorpb.ComputeAverageRequest{
			Number: 13,
		},
		calculatorpb.ComputeAverageRequest{
			Number: 21,
		},
	}

	for _, r := range request {
		err := s.Send(&r)

		if err == io.EOF {
			log.Fatalf("Server close the request : %+v \n", err)
		}

		if err != nil {
			log.Fatalf("Error while sending request to compute average : %+v \n", err)
		}
	}
	res, err := s.CloseAndRecv()

	if err != nil {
		log.Fatalf("Server sent error response : %+v \n", err)
	}

	fmt.Println("=======================================================================")
	fmt.Println(res)
	fmt.Println("=======================================================================")
}
func doUnary(c calculatorpb.CalculatorServiceClient) {
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	strm, err := c.Prime(context.Background(), &calculatorpb.CalcPrimeRequest{
		Number: 729653,
	})

	if err != nil {
		log.Fatalf("Error occurred : %s\n", err.Error())
	}

	for {
		strmRes, err := strm.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error occurred on stearming client : %s", err.Error())
		}

		fmt.Println("=======================================================================")
		fmt.Println(strmRes.Result)
		fmt.Println("=======================================================================")
	}
}
