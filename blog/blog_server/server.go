package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/dilrandi/grpc-sample-app/blog/blogpb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type blogServer struct {
}
type Person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	// if we crash, will get files name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog server ... !")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		log.Fatalf("Error while creating mongo client : %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatalf("Error while connecting to mongo server : %v", err)
	}

	res, err := client.Database("deleema").Collection("HelloWorld").InsertOne(ctx,
		&Person{
			Age:   28,
			Fname: "Deleema",
			Lname: "Fernando",
		})

	if err != nil {
		log.Fatalf("Errorrr : %v", err)
	}

	fmt.Println("=======================================================================")
	fmt.Printf("\n\n %+v \n\n", res)
	fmt.Println("=======================================================================")

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
