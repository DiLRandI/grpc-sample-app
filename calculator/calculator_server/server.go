package main

import (
	"context"
	"fmt"
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
func main() {
	fmt.Println("Calculator server starting")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
