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

	//Add the user to the map
	server.currentUsers[tmp] = nil
	lamportTime = compareLamportTime(int(req.TimeStamp))
	lamportTime++
	srLock.Unlock()

	log.Printf(req.User.Name + " joined at Lamport Time: " + fmt.Sprint(lamportTime))

	var msg string = req.User.Name + " joined the chat at Lamport Time: "

	for i, user := range server.currentUsers {
		if i != req.User.UserID {
			srLock.Lock()
			lamportTime++
			srLock.Unlock()
			user.Send(&proto.PostResponse{
				User:      req.User,
				Message:   msg,
				TimeStamp: int32(lamportTime),
			})
			log.Printf("broadcasted join message at Lamport Time: %d to user %d", lamportTime, i)
		}
	}

	srLock.Lock()
	lamportTime++
	srLock.Unlock()

	joinResponse := &proto.JoinResponse{
		UserID:    tmp,
		TimeStamp: int32(lamportTime),
		Message:   msg,
	}

	log.Printf("server LT after JoinResponse %d", lamportTime)

	return joinResponse, nil
}

func (server *ChittyChatServiceServer) LeaveServer(ctx context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {
	srLock.Lock()
	delete(server.currentUsers, req.User.UserID)
	srLock.Unlock()

	srLock.Lock()
	lamportTime = compareLamportTime(int(req.TimeStamp))
	lamportTime++
	srLock.Unlock()
	log.Printf(req.User.Name + " left at Lamport Time: " + fmt.Sprint(lamportTime))

	var msg string = req.User.Name + " left the chat at Lamport Time: "

	srLock.Lock()
	for i, user := range server.currentUsers {
		if i != req.User.UserID {
			lamportTime++
			user.Send(&proto.PostResponse{
				User:      req.User,
				Message:   msg,
				TimeStamp: int32(lamportTime),
			})
		}
	}
	srLock.Unlock()

	lamportTime = compareLamportTime(int(req.TimeStamp))
	lamportTime++

	leaveResponse := &proto.LeaveResponse{
		Message:   "Goodbye " + req.User.Name + ", we hope to see you again soon!!",
		TimeStamp: int32(lamportTime),
	}

	return leaveResponse, nil
}

func (server *ChittyChatServiceServer) Communicate(stream proto.ChittyChatService_CommunicateServer) error {
	srLock.Lock()
	//Links the stream to the user
	server.currentUsers[int32(server.totalAmountUsers-1)] = stream

	lamportTime = lamportTime + 3 //since we cannot send a request in the same way as the other methods,
	//then we increment to make the servers lamportTime the same as the clients (+2) and then increment once more because a rpc call has been made(+1)

	log.Printf("New user established communication")
	srLock.Unlock()

	//Sends the messages to the different users
	for {
		message, err := stream.Recv()
		if message != nil {

			srLock.Lock()
			lamportTime = compareLamportTime(int(message.TimeStamp))
			lamportTime++
			log.Printf("Server TimeStamp: %d", lamportTime)
			log.Printf("Message timeStamp: %d", message.TimeStamp)

			srLock.Unlock()

			if err != nil {
				log.Printf("User has disconnected")

			} else {
				var msg string = message.Message + " ----- server received at Lamport Time: "
				log.Printf(msg + fmt.Sprint(lamportTime))

				//Broadcast the message to all users
				srLock.Lock()
				for i, user := range server.currentUsers {
					if i != message.User.UserID {
						lamportTime++
						user.Send(&proto.PostResponse{
							Message:   message.Message + " ----- client received at Lamport Time: ",
							User:      message.User,
							TimeStamp: int32(lamportTime),
						})
					}
				}
				srLock.Unlock()
			}
		}

	}

}

func compareLamportTime(otherLamportTime int) int {
	if lamportTime > otherLamportTime {
		return lamportTime
	} else if otherLamportTime > lamportTime {
		return otherLamportTime
	} else {
		return lamportTime
	}
}
