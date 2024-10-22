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

func (server *ChittyChatServiceServer) PublishMessage(ctx context.Context, req *proto.PostMessage) (*proto.PostResponse, error) {
	log.Printf(req.User.Name + " published a message")

	//Checks for the validity of the message, requesting to be posted.
	if !utf8.ValidString(req.Message) {
		server.lamportTime += 1
		return &proto.PostResponse{
			Message: "Messages must be valid in utf8",
		}, nil
	} else if len(req.Message) > 128 {
		server.lamportTime += 1
		return &proto.PostResponse{
			Message: "Messages cant be more than 128 characters long",
		}, nil
	} else if len(req.Message) <= 0 {
		server.lamportTime += 1
		return &proto.PostResponse{
			Message: "Messages must be longer than 0 characters",
		}, nil
	}
	/*if strings.Contains(req.Message, " joined at Lamport Time: ") {
		server.currentUsers = append(server.currentUsers, req.User)
		req.User.UserID = server.totalAmuntUsers
		server.totalAmuntUsers++
	}*/
	server.messages = append(server.messages, *req)

	log.Printf("Message: %s", req.Message)

	server.lamportTime += 1

	return &proto.PostResponse{
		Message: "Message published successfully",
	}, nil
}

func (server *ChittyChatServiceServer) compareLT(otherLT int) int {
	//Compares the servers lamport Timestamp with a clients
	var result int

	if server.lamportTime > int(otherLT) {
		result = server.lamportTime
	} else {
		result = int(otherLT)
	}

	return result
}

func (server *ChittyChatServiceServer) BroadcastMessages(req *proto.BroadcastRequest, stream proto.ChittyChatService_BroadcastMessagesServer) error {
	log.Printf(req.User.Name + " broadcasts all messages")
	//Needs to be a pointer, since we in our methods send pointers
	//We send the entire array over, instead of each messsage idividually.
	var tmpMessages []*proto.PostMessage
	for i := range server.messages {
		tmpMessages = append(tmpMessages, &proto.PostMessage{
			User:      server.messages[i].User,
			Message:   server.messages[i].Message,
			TimeStamp: server.messages[i].TimeStamp,
		})
	}

	for i := range server.currentUsers {
		server.currentUsers[i].Send(&proto.BroadcastResponse{
			Messages:  tmpMessages,
			TimeStamp: int32(server.lamportTime),
		})
	}

	return nil
}
