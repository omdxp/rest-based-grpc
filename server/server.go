package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/Omar-Belghaouti/rest-based-grpc/gen/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		// http server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTestApiServer(s, &testApiServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
