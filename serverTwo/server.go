package main

import (
	"context"
	proto "handin3/grpc"
	"log"
	"net"
	"sync"
	"unicode/utf8"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	messages []proto.PostMessage

	currentUsers []proto.ChittyChatService_BroadcastMessagesServer

	totalAmuntUsers int32

	lamportTime int
}

var srLock sync.Mutex

func main() {
	server := &ChittyChatServiceServer{messages: []proto.PostMessage{}, totalAmuntUsers: 0, lamportTime: 0}

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

func(server *ChittyChatServiceServer) joinServer(req *proto.Empty) (*proto.JoinResponse){	
	var tmp = server.totalAmuntUsers
	server.totalAmuntUsers++
	return &proto.JoinResponse{
		UserID: tmp,
	}
}

func(server *ChittyChatServiceServer) connected(stream *proto.ChittyChatService_ConnectedServer) (*proto.ChittyChatService_ConnectedServer) {
	
	go messages(){
		for{
			message, erro := stream.Recv()
			if (erro==nill){
				
			}
		}

	}
}