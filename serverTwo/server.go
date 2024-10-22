package main

import (
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

	//This starts the server
	server.start_server()

}

func (server *ChittyChatServiceServer) start_server() {
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

func (server *ChittyChatServiceServer) joinServer(req *proto.Empty) *proto.JoinResponse {
	var tmp = server.totalAmountUsers
	server.totalAmountUsers++
	return &proto.JoinResponse{
		UserID: tmp,
	}
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
}
