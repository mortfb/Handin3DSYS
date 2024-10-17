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

type clientStreams struct {
	broadCastAll   proto.ChittyChatService_BroadcastAllMessagesServer
	broadCastJoin  proto.ChittyChatService_BroadcastJoinServer
	broadCastLeave proto.ChittyChatService_BroadcastLeaveServer
}

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	messages []proto.PostMessage

	activeUsers map[int32]clientStreams

	lambortTime int
	muLock      sync.Mutex
}

var totalAmuntUsers int = 0

func main() {
	server := &ChittyChatServiceServer{messages: []proto.PostMessage{}, activeUsers: map[int32]clientStreams{}, lambortTime: 0}

	//This starts the server
	server.start_server()
}

func (s *ChittyChatServiceServer) start_server() {
	log.Printf("Server started")
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

func (s *ChittyChatServiceServer) PublishMessage(ctx context.Context, req *proto.PostMessage) (*proto.PostResponse, error) {
	log.Printf(req.User.Name + " published a message")
	s.muLock.Lock()
	s.lambortTime += 1
	s.lambortTime = s.compareLT(int(req.TimeStamp))
	req.TimeStamp = int32(s.lambortTime)
	s.muLock.Unlock()

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

	s.messages = append(s.messages, *req)

	log.Printf("Message: %s", req.Message)

	for i := range s.messages {
		log.Printf("Message: %s %s", s.messages[i].User.Name, s.messages[i].Message)
	}

	s.muLock.Lock()
	s.lambortTime += 1
	s.muLock.Unlock()

	return &proto.PostResponse{
		Success: true,
		Message: "Message published successfully",
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
	log.Printf(req.User.Name + " broadcasts all messages")
	s.muLock.Lock()
	check, exists := s.activeUsers[req.User.UserID]
	if exists {
		//if the user already exists, we update the stream
		check.broadCastAll = stream
		s.activeUsers[req.User.UserID] = check
	} else {
		s.activeUsers[req.User.UserID] = clientStreams{
			broadCastAll: stream}
	}
	s.muLock.Unlock()

	s.muLock.Lock()
	s.lambortTime += 1
	s.lambortTime = s.compareLT(int(req.TimeStamp))
	s.muLock.Unlock()

	//sets up a timer, that executes at a certain interval
	timer := time.NewTicker(4 * time.Second)

	for {
		select {
		case <-stream.Context().Done():
			return nil

		case <-timer.C:
			for _, user := range s.activeUsers {
				for i := range s.messages {
					message := &s.messages[i]
					err := user.broadCastAll.Send(&proto.BroadcastAllResponse{
						Messages:  message,
						TimeStamp: int32(s.lambortTime),
					})
					if err != nil {
						log.Println(err.Error())
						return err
					}
				}
			}
		}
	}
}

func (s *ChittyChatServiceServer) BroadcastJoin(req *proto.NewClientJoinedRequest, stream proto.ChittyChatService_BroadcastJoinServer) error {
	log.Printf(req.User.Name + " joins the chat")
	s.muLock.Lock()
	check, exists := s.activeUsers[req.User.UserID]
	if exists {
		//if the user already exists, we update the stream
		check.broadCastJoin = stream
		s.activeUsers[req.User.UserID] = check
	} else {
		s.activeUsers[req.User.UserID] = clientStreams{
			broadCastJoin: stream}
	}
	s.muLock.Unlock()

	s.muLock.Lock()
	s.lambortTime += 1
	s.lambortTime = s.compareLT(int(req.TimeStamp))
	s.muLock.Unlock()

	//Server updates the user ID
	s.muLock.Lock()
	if totalAmuntUsers == 0 {
		totalAmuntUsers += 1
		req.User.UserID = 0
	} else {
		req.User.UserID = int32(totalAmuntUsers)
		totalAmuntUsers += 1
	}
	s.muLock.Unlock()

	var message = "User " + req.User.Name + " joined at Lamport Time " + strconv.Itoa(s.lambortTime)

	for _, user := range s.activeUsers {
		err := user.broadCastJoin.Send(&proto.NewClientJoinedResponse{
			Message:   message,
			TimeStamp: int32(s.lambortTime),
		})
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (s *ChittyChatServiceServer) BroadcastLeave(req *proto.ClientLeaveRequest, stream proto.ChittyChatService_BroadcastLeaveServer) error {
	log.Printf(req.User.Name + " leaves")
	s.muLock.Lock()
	check, exists := s.activeUsers[req.User.UserID]
	if exists {
		//if the user already exists, we update the stream
		check.broadCastLeave = stream
		s.activeUsers[req.User.UserID] = check
	} else {
		s.activeUsers[req.User.UserID] = clientStreams{
			broadCastLeave: stream}
	}
	s.muLock.Unlock()

	s.muLock.Lock()
	s.lambortTime += 1
	s.lambortTime = s.compareLT(int(req.TimeStamp))
	s.muLock.Unlock()

	var message = "User " + req.User.Name + " left at Lamport Time " + strconv.Itoa(s.lambortTime)

	for _, user := range s.activeUsers {
		err := user.broadCastLeave.Send(&proto.ClientLeaveResponse{
			Message:   message,
			TimeStamp: int32(s.lambortTime),
		})
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	delete(s.activeUsers, req.User.UserID)

	return nil
}

//May need to go back to a select case
/*
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
*/
