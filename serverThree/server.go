package main

import (
	"fmt"

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

	srLock.Lock()
	server.lamportTime = compareLamportTime(server.lamportTime, int(req.TimeStamp))
	srLock.Unlock()

	log.Printf(req.User.Name + " joined at Lamport Time: " + fmt.Sprint(server.lamportTime))
	server.currentUsers[tmp] = nil

	log.Printf("broadcasted join message")
	var msg string = req.User.Name + " joined the chat at Lamport Time: "

	message := &proto.PostMessage{
		User:      req.User,
		Message:   msg,
		TimeStamp: int32(server.lamportTime),
	}

	server.BroadcastMessage(message)

	srLock.Lock()
	server.lamportTime = compareLamportTime(server.lamportTime, int(req.TimeStamp))
	srLock.Unlock()

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(server.lamportTime),
		Message:   "Welcome to the server " + req.User.Name,
	}

	return joinResponse, nil
}

func (server *ChittyChatServiceServer) BroadcastMessage(message *proto.PostMessage) error {
	for i, user := range server.currentUsers {
		if i != message.User.UserID {
			if user != nil {
				srLock.Lock()
				server.lamportTime = compareLamportTime(server.lamportTime, int(message.TimeStamp))
				srLock.Unlock()
				user.Send(message)
			}
		}
	}
	return nil
}

func (server *ChittyChatServiceServer) LeaveServer(ctx context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {
	srLock.Lock()
	delete(server.currentUsers, req.User.UserID)
	srLock.Unlock()

	srLock.Lock()
	server.lamportTime = compareLamportTime(server.lamportTime, int(req.TimeStamp))
	srLock.Unlock()
	log.Printf(req.User.Name + " left at Lamport Time: " + fmt.Sprint(server.lamportTime))

	leaveResponse := &proto.LeaveResponse{
		Message:   "Goodbye " + req.User.Name + ", we hope to see you again soon!",
		TimeStamp: int32(server.lamportTime),
	}

	var msg string = req.User.Name + " left the chat at Lamport Time: "
	message := &proto.PostMessage{
		User:      req.User,
		Message:   msg,
		TimeStamp: int32(server.lamportTime),
	}

	server.BroadcastMessage(message)
	server.lamportTime = compareLamportTime(server.lamportTime, int(req.TimeStamp))

	return leaveResponse, nil
}

func (server *ChittyChatServiceServer) Communicate(stream proto.ChittyChatService_CommunicateServer) error {
	srLock.Lock()
	server.currentUsers[int32(server.totalAmountUsers-1)] = stream
	srLock.Unlock()

	//Shoud send the messages to the different users
	for {
		message, err := stream.Recv()
		if err == nil {
			srLock.Lock()
			server.lamportTime = compareLamportTime(int(message.TimeStamp), server.lamportTime)
			srLock.Unlock()
			var msg string = message.Message + " ----- sends at Lamport Time: " //+ fmt.Sprint(server.lamportTime)

			BroadcastMessage := &proto.PostMessage{
				Message:   msg,
				User:      message.User,
				TimeStamp: int32(server.lamportTime),
			}

			log.Printf(BroadcastMessage.Message + fmt.Sprint(server.lamportTime))

			server.BroadcastMessage(BroadcastMessage)
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
