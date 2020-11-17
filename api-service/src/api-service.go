package main

import (
	"fmt"
	"context"
    "log"
    "net"
    "./proto/messagepb"
    "google.golang.org/grpc"
)

type SimpleServer struct{
	messagepb.UnimplementedMessageServiceServer
}

func (*SimpleServer) Message(ctx context.Context, req *messagepb.MessageRequest) (*messagepb.MessageResponse, error) {
	reqMessage := req.GetMessage()
	fmt.Println("req: ", reqMessage)
	resMessage := "hello"
	return &messagepb.MessageResponse{
		Message: resMessage,
	}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

	srv := grpc.NewServer()
    messagepb.RegisterMessageServiceServer(srv, &SimpleServer{})

    if err := srv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}