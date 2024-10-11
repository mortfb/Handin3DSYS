package main

import (
	"context"
	proto "handin3/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	messages []proto.PostMessage //maybe not needed

	lambortTime int

	muLock sync.Mutex
}

func main() {
	server := &ChittyChatServiceServer{messages: []proto.PostMessage{}, lambortTime: 0}
	/*
		//This is just to test the append
		server.messages = append(server.messages, proto.PublishRequest{
			User:      "test",
			Message:   "test",
			TimeStamp: 0,
		})
	*/

	//This starts the server
	server.start_server()

	//Makes the server do the actions we need
	//server.run_server()

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

func (s *ChittyChatServiceServer) run_server() {

}

func (s *ChittyChatServiceServer) PublishMessage(ctx context.Context, req *proto.PostMessage) (*proto.PostResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))
	req.TimeStamp = int32(s.lambortTime)

	if !utf8.ValidString(req.Message) {
		s.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages must be valid in utf8",
		}, nil
	} else if len(req.Message) > 128 {
		s.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages cant be more than 128 characters long",
		}, nil
	} else if len(req.Message) <= 0 {
		s.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages must be longer than 0 characters",
		}, nil
	}

	//If it is valid
	s.messages = append(s.messages, *req)

	log.Printf("Message: %s", req.Message)

	for i := range s.messages {
		log.Printf("Message: %s", s.messages[i].User.Name, " ", s.messages[i].Message)
	}

	s.lambortTime += 1
	return &proto.PostResponse{
		Success: true,
		Message: "Message published successfully",
	}, nil
}

func (s *ChittyChatServiceServer) NewClientJoined(ctx context.Context, req *proto.NewClientJoinedRequest) (*proto.NewClientJoinedResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))

	s.lambortTime += 1
	return &proto.NewClientJoinedResponse{
		Message:   req.User.Name + " joined successfully at Lamport time " + strconv.Itoa(s.lambortTime),
		TimeStamp: int32(s.lambortTime),
	}, nil
}

func (s *ChittyChatServiceServer) ClientLeave(ctx context.Context, req *proto.ClientLeaveRequest) (*proto.ClientLeaveResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))

	//Receives and sends, therefore s.lambortTime X 2
	s.lambortTime += 1
	return &proto.ClientLeaveResponse{
		Message:   req.User.Name + " left at Lamport time " + strconv.Itoa(s.lambortTime), //req.Message is the name of the client
		TimeStamp: int32(s.lambortTime),
	}, nil
}

func (s *ChittyChatServiceServer) compareLT(otherLT int) int {
	var result int

	if s.lambortTime > int(otherLT) {
		result = s.lambortTime
	} else {
		result = int(otherLT)
	}

	return result
}

func (s *ChittyChatServiceServer) BroadcastAllMessages(req *proto.BroadcastAllRequest, stream proto.ChittyChatService_BroadcastAllMessagesServer) error {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))
	//sets up a timer, that executes every 3 seconds
	timer := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-stream.Context().Done():
			return nil

		case <-timer.C:
			for _, message := range s.messages {
				err := stream.Send(&proto.BroadcastAllResponse{
					Messages: &message,
				})
				if err != nil {
					log.Println(err.Error())
					return err
				}
			}
		}
	}
}
