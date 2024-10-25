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
}

var srLock sync.Mutex
var lamportTime int = 0

func main() {
	server := &ChittyChatServiceServer{
		currentUsers:     make(map[int32]proto.ChittyChatService_CommunicateServer),
		totalAmountUsers: 0,
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

	server.currentUsers[tmp] = nil
	fmt.Print(req.User.Name + string(req.TimeStamp))
	lamportTime = compareLamportTime(int(req.TimeStamp))
	srLock.Unlock()

	log.Printf(req.User.Name + " joined at Lamport Time: " + fmt.Sprint(lamportTime))

	log.Printf("broadcasted join message")
	var msg string = req.User.Name + " joined the chat at Lamport Time: "

	message := &proto.PostMessage{
		User:      req.User,
		Message:   msg,
		TimeStamp: int32(lamportTime),
	}

	server.BroadcastMessage(message)

	srLock.Lock()
	//Dont need to compare lamport time here
	lamportTime++
	srLock.Unlock()

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(lamportTime),
		Message:   "Welcome to the server " + req.User.Name,
	}

	return joinResponse, nil
}

func (server *ChittyChatServiceServer) BroadcastMessage(message *proto.PostMessage) {
	srLock.Lock()
	defer srLock.Unlock()
	for i, user := range server.currentUsers {
		if i != message.User.UserID {
			//if user != nil {
			lamportTime++
			user.Send(message)
			//}
		}
	}
}

func (server *ChittyChatServiceServer) LeaveServer(ctx context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {
	srLock.Lock()
	delete(server.currentUsers, req.User.UserID)
	srLock.Unlock()

	srLock.Lock()
	lamportTime = compareLamportTime(int(req.TimeStamp))
	srLock.Unlock()
	log.Printf(req.User.Name + " left at Lamport Time: " + fmt.Sprint(lamportTime))

	leaveResponse := &proto.LeaveResponse{
		Message:   "Goodbye " + req.User.Name + ", we hope to see you again soon!",
		TimeStamp: int32(lamportTime),
	}

	var msg string = req.User.Name + " left the chat at Lamport Time: "
	message := &proto.PostMessage{
		User:      req.User,
		Message:   msg,
		TimeStamp: int32(lamportTime),
	}

	server.BroadcastMessage(message)
	lamportTime = compareLamportTime(int(req.TimeStamp))

	return leaveResponse, nil
}

func (server *ChittyChatServiceServer) Communicate(stream proto.ChittyChatService_CommunicateServer) error {
	srLock.Lock()
	server.currentUsers[int32(server.totalAmountUsers-1)] = stream
	srLock.Unlock()

	log.Printf("User connected")
	//Should send the messages to the different users
	for {
		message, err := stream.Recv()
		srLock.Lock()
		lamportTime = compareLamportTime(int(message.TimeStamp))
		srLock.Unlock()

		if err != nil {
			log.Printf("User disconnected")

		} else {
			var msg string = message.Message + " ----- sent at Lamport Time: " //+ fmt.Sprint(server.lamportTime)

			BroadcastMessage := &proto.PostMessage{
				Message:   msg,
				User:      message.User,
				TimeStamp: int32(lamportTime),
			}
			server.BroadcastMessage(BroadcastMessage)

			log.Printf(BroadcastMessage.Message + fmt.Sprint(lamportTime))
		}

	}

}

func compareLamportTime(otherLamportTime int) int {
	if lamportTime > otherLamportTime {
		return lamportTime + 1
	} else if otherLamportTime > lamportTime {
		return otherLamportTime + 1
	} else {
		return lamportTime + 1
	}
}
