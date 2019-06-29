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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type blogServer struct {
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
	AuthorID string             `bson:"author_id"`
}

var collection *mongo.Collection

func main() {
	// if we crash, will get files name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog server ... !")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))

	if err != nil {
		log.Fatalf("Error while creating mongo client : %v", err)
	}

	ctx, cf := context.WithTimeout(context.Background(), 10*time.Second)
	defer cf()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatalf("Error while connecting to mongo server : %v", err)
	}

	collection = client.Database("myBlogsDb").Collection("blogs")

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
	fmt.Println("Disconnecting from mongo server")
	client.Disconnect(ctx)
	fmt.Println("Closing the server application")

}
