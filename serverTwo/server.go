package main

import (
	"context"
	"fmt"
	proto "handin3/grpc"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer

	currentUsers []proto.ChittyChatService_ConnectedServer

	totalAmountUsers int32

	lamportTime int
}

var srLock sync.Mutex

func main() {
	server := &ChittyChatServiceServer{
		currentUsers:     []proto.ChittyChatService_ConnectedServer{},
		totalAmountUsers: 0,
		lamportTime:      0,
	}

	log.Printf("Server started")
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Did not work")
	}

	//Register the server, from proto-file
	proto.RegisterChittyChatServiceServer(grpcServer, server)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Did not work")
	}

}

func (server *ChittyChatServiceServer) JoinServer(ctx context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
	var tmp = server.totalAmountUsers
	server.totalAmountUsers++

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(server.lamportTime),
	}

	joinMessage := &proto.PostMessage{
		Message:   fmt.Sprintf("User %s has joined the chat", req.User.Name),
		User:      req.User,
		TimeStamp: int32(server.lamportTime),
	}

	for _, client := range server.currentUsers {
		if joinMessage.User.UserID != tmp {
			err := client.Send(joinMessage)
			if err != nil {
				fmt.Println("cant print the joinMessage")
			}

		}

	}

	return joinResponse, nil

}

func (server *ChittyChatServiceServer) messages(stream proto.ChittyChatService_ConnectedServer) {
	for {
		message, err := stream.Recv()
		if err == nil {
			for i := range server.currentUsers {
				if i != int(message.User.UserID) { //userID
					server.currentUsers[i].Send(message)
				}
			}
		}
	}
}

func (server *ChittyChatServiceServer) connected(req *proto.PostMessage, stream proto.ChittyChatService_ConnectedServer) *proto.ChittyChatService_ConnectedServer {
	server.currentUsers = append(server.currentUsers, stream)
	go server.messages(stream)

	return &stream
}
