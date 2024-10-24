package main

import (
	proto "Handin3DSYS/grpc"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer

	currentUsers map[int32]proto.ChittyChatService_CommunicateServer

	totalAmountUsers int32

	lamportTime int
}

var srLock sync.Mutex

func main() {
	server := &ChittyChatServiceServer{
		currentUsers:     make(map[int32]proto.ChittyChatService_CommunicateServer),
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
	srLock.Lock()
	var tmp = server.totalAmountUsers
	req.User.UserID = tmp
	server.totalAmountUsers++
	srLock.Unlock()

	log.Println("User joined: " + req.User.Name)

	server.lamportTime++

	server.currentUsers[tmp] = nil

	log.Println("broadcasted join message")

	var message string = "User " + req.User.Name + " joined the chat at: " + string(server.lamportTime)
	server.BroadcastMessage(message, req.User)

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(server.lamportTime),
		Message:   "Welcome to the server" + req.User.Name,
	}

	return joinResponse, nil
}

func (server *ChittyChatServiceServer) BroadcastMessage(message string, client *proto.User) error {
	for i, user := range server.currentUsers {
		if i != client.UserID {
			if user != nil {
				user.Send(&proto.PostMessage{
					Message:   message,
					TimeStamp: int32(server.lamportTime),
				})
			}
		}
	}
	return nil
}

func (server *ChittyChatServiceServer) LeaveServer(ctx context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {
	srLock.Lock()
	delete(server.currentUsers, req.User.UserID)
	srLock.Unlock()

	log.Println("User left: " + req.User.Name)

	leaveResponse := &proto.LeaveResponse{
		Message:   "Goodbye " + req.User.Name + ", we hope to see you again soon!",
		TimeStamp: int32(server.lamportTime),
	}

	var message string = "User left: " + req.User.Name
	server.BroadcastMessage(message, req.User)

	return leaveResponse, nil
}

func (server *ChittyChatServiceServer) Communicate(stream proto.ChittyChatService_CommunicateServer) error {

	//Shoud send the messages to the different users
	for {
		message, err := stream.Recv()
		if err == nil {
			server.lamportTime = compareLamportTime(int(message.TimeStamp), server.lamportTime)
			server.currentUsers[message.User.UserID] = stream

			var BroadcastMessage string = message.User.Name + ": " + message.Message + " at time: " + string(server.lamportTime)
			log.Println(BroadcastMessage)

			server.BroadcastMessage(BroadcastMessage, message.User)
		}

	}

}

func compareLamportTime(lamportTime int, messageLamportTime int) int {
	if lamportTime > messageLamportTime {
		return lamportTime + 1
	} else {
		lamportTime = messageLamportTime
		return lamportTime + 1
	}
}
