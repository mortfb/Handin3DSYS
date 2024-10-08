package main

import (
	proto "handin3/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	users []string //maybe not needed
}

func main() {
	server := &ChittyChatServiceServer{users: []string{}}
	//This starts the server
	server.start_server()

}

func (s *ChittyChatServiceServer) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Did not work")
	}

	//Register the server, from proto-file
	proto.RegisterChittyChatServiceServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Did not work")
	}
}
