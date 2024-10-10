package main

import (
	"context"
	proto "handin3/grpc"
	"log"
	"net"
	"strconv"
	"unicode/utf8"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	messages    []proto.PublishRequest //maybe not needed
	lambortTime int
}

func main() {
	server := &ChittyChatServiceServer{messages: []proto.PublishRequest{}, lambortTime: 0}

	//This is just to test the append
	server.messages = append(server.messages, proto.PublishRequest{
		User:      "test",
		Message:   "test",
		TimeStamp: 0,
	})

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

func (s *ChittyChatServiceServer) PublishMessage(ctx context.Context, req *proto.PublishRequest) (*proto.PublishResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))
	req.TimeStamp = int32(s.lambortTime)

	if !utf8.ValidString(req.Message) {
		return &proto.PublishResponse{
			Success: false,
			Message: "Messages must be valid in utf8",
		}, nil
	} else if len(req.Message) > 128 {
		return &proto.PublishResponse{
			Success: false,
			Message: "Messages cant be more than 128 characters long",
		}, nil
	} else if len(req.Message) <= 0 {
		return &proto.PublishResponse{
			Success: false,
			Message: "Messages must be longer than 0 characters",
		}, nil
	}

	//If it is valid
	s.messages = append(s.messages, *req)

	log.Printf("Message: %s", req.Message)

	for i := range s.messages {
		log.Printf("Message: %s", s.messages[i].Message)
	}

	return &proto.PublishResponse{
		Success: true,
		Message: "Message published successfully",
	}, nil
}

func (s *ChittyChatServiceServer) NewClientJoined(ctx context.Context, req *proto.NewClientJoinedRequest) (*proto.NewClientJoinedResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))

	return &proto.NewClientJoinedResponse{
		Name:      req.Name + " joined successfully at Lamport time " + strconv.Itoa(s.lambortTime),
		TimeStamp: int32(s.lambortTime),
	}, nil
}

func (s *ChittyChatServiceServer) ClientLeave(ctx context.Context, req *proto.ClientLeaveRequest) (*proto.ClientLeaveResponse, error) {
	s.lambortTime += 1

	s.lambortTime = s.compareLT(int(req.TimeStamp))

	return &proto.ClientLeaveResponse{
		Name:      req.Name + " left at Lamport time " + strconv.Itoa(s.lambortTime), //req.Message is the name of the client
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

/*
func (s *ChittyChatServiceServer) BroadcastAll(in *proto.BroadcastAllRequest, stream proto.ChittyChatService_BroadcastAllServer) (*proto.BroadcastAllresponse, error) {
	return nil
}

*/
