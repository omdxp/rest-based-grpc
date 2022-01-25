package main

import (
	"context"
	"log"

	pb "github.com/Omar-Belghaouti/rest-based-grpc/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := pb.NewTestApiClient(conn)

	res, err := client.Echo(context.Background(), &pb.ResponseRequest{
		Msg: "Hello Everybody",
	})
	if err != nil {
		log.Fatalf("failed to call Echo: %v", err)
	}
	log.Printf("Echo response: %v", res.Msg)

}
