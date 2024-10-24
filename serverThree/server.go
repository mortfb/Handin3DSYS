package main

import (
	proto "Handin3DSYS/grpc"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer

	currentUsers map[int32]proto.ChittyChatService_ConnectedServer

	totalAmountUsers int32

	lamportTime int
}

var srLock sync.Mutex

func main() {
	server := &ChittyChatServiceServer{
		currentUsers:     make(map[int32]proto.ChittyChatService_ConnectedServer),
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

func (server *ChittyChatServiceServer) JoinServer(req *proto.JoinRequest, stream proto.ChittyChatService_JoinServerServer) error {
	srLock.Lock()
	var tmp = server.totalAmountUsers
	req.User.UserID = tmp
	server.totalAmountUsers++
	srLock.Unlock()

	log.Println("User joined: " + req.User.Name)

	server.currentUsers[tmp] = nil

	for _, user := range server.currentUsers {
		if user != nil {
			user.Send(&proto.PostMessage{
				Message:   fmt.Sprintf("User %s has joined the chat", req.User.Name),
				TimeStamp: int32(server.lamportTime),
			})
		}
	}

	log.Println("broadcasted join message")

	var message string = "New user joined: " + req.User.Name
	server.BroadcastMessage(message, req.User)

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(server.lamportTime),
		Message:   "Welcome to the server" + req.User.Name,
	}

	stream.Send(joinResponse)

	return nil
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

func (server *ChittyChatServiceServer) LeaveServer(req *proto.LeaveRequest, stream proto.ChittyChatService_LeaveServerServer) error {
	srLock.Lock()
	delete(server.currentUsers, req.User.UserID)
	srLock.Unlock()

	stream.Send(&proto.LeaveResponse{
		Message:   "Goodbye " + req.User.Name + ", we hope to see you again soon!",
		TimeStamp: int32(server.lamportTime),
	})

	var message string = "User left: " + req.User.Name
	server.BroadcastMessage(message, req.User)

	return nil
}

func (server *ChittyChatServiceServer) Connected(stream proto.ChittyChatService_ConnectedServer) error {

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
