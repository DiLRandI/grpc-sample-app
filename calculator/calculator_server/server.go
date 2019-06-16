package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dilrandi/grpc-sample-app/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, r *calculatorpb.CalcRequest) (*calculatorpb.CalcResponse, error) {
	fmt.Println("=======================================================================")
	fmt.Println(r, ctx)
	fmt.Println("=======================================================================")
	res := &calculatorpb.CalcResponse{
		Result: r.No1 + r.No2,
	}
	return res, nil
}

func (*server) Prime(r *calculatorpb.CalcPrimeRequest, s calculatorpb.CalculatorService_PrimeServer) error {
	num := r.Number
	var dev int32 = 2

	for num > 1 {
		if num%dev == 0 {
			s.Send(&calculatorpb.CalcResponse{
				Result: dev,
			})
			num = num / dev
		} else {
			dev++
		}

		fmt.Println("=======================================================================")
		fmt.Println(num, dev)
		fmt.Println("=======================================================================")

		// time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func main() {
	fmt.Println("Calculator server starting")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error occurred : %s \n", err.Error())
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
