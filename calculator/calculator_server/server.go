package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) ComputeAverage(s calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Computing average ")
	var n int32
	var tot int32

	for {
		m, err := s.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf(" Error while computing avg : %s \n", err.Error())
		}

		tot += m.Number
		n++
	}

	fmt.Printf("Total : %v and Elements : %v \n", tot, n)
	avg := float32(tot) / float32(n)
	s.SendAndClose(&calculatorpb.ComputeAverageResponse{
		Result: avg,
	})

	return nil
}

func (*server) FindMax(strm calculatorpb.CalculatorService_FindMaxServer) error {
	fmt.Println("Finding max")
	var max int32
	for {
		r, err := strm.Recv()
		if err == io.EOF {
			fmt.Println("Client closed the stream!")
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving data from client %+v", err)
		}
		fmt.Println("=======================================================================")
		fmt.Printf("Comparing max %d with received value %d \n", max, r.Number)
		fmt.Println("=======================================================================")
		if r.Number >= max {
			max = r.Number
			strm.Send(&calculatorpb.FindMaxResponse{
				Result: max,
			})
		}
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
