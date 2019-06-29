package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/dilrandi/grpc-sample-app/blog/blogpb"

	"google.golang.org/grpc"
)

type blogServer struct {
}

func main() {
	// if we crash, will get files name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog server ... !")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Error while string the Blog server : %v", err)
	}

	opts := []grpc.ServerOption{}
	svr := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(svr, &blogServer{})

	go func() {
		fmt.Println("Starting to serve ...")
		if err := svr.Serve(lis); err != nil {
			log.Fatalf("Error while serving : %v", err)
		}
	}()

	// wating for Control + C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// wait till chanel receive some data
	<-ch
	fmt.Println("Stopping the server")
	svr.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing the server application")
}
