package main

import (
	"context"
	proto "handin3/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"unicode/utf8"

	"google.golang.org/grpc"
)

// Struct containing all the streams, we have
type streamsForClients struct {
	broadCastAll   proto.ChittyChatService_BroadcastAllMessagesServer
	broadCastJoin  proto.ChittyChatService_BroadcastJoinServer
	broadCastLeave proto.ChittyChatService_BroadcastLeaveServer
}

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer
	messages []proto.PostMessage

	currentUsers map[int32]streamsForClients

	lambortTime int
}

var srLock sync.Mutex

var totalAmuntUsers int = 0

func main() {
	server := &ChittyChatServiceServer{messages: []proto.PostMessage{}, currentUsers: map[int32]streamsForClients{}, lambortTime: 0}

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
		server.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages must be valid in utf8",
		}, nil
	} else if len(req.Message) > 128 {
		server.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages cant be more than 128 characters long",
		}, nil
	} else if len(req.Message) <= 0 {
		server.lambortTime += 1
		return &proto.PostResponse{
			Success: false,
			Message: "Messages must be longer than 0 characters",
		}, nil
	}

	server.messages = append(server.messages, *req)

	log.Printf("Message: %s", req.Message)

	return &proto.PostResponse{
		Success: true,
		Message: "Message published successfully",
	}, nil
}

func (server *ChittyChatServiceServer) compareLT(otherLT int) int {
	//Compares the servers lamport Timestamp with a clients
	var result int

	if server.lambortTime > int(otherLT) {
		result = server.lambortTime
	} else {
		result = int(otherLT)
	}

	return result
}

func (server *ChittyChatServiceServer) BroadcastAllMessages(req *proto.BroadcastAllRequest, stream proto.ChittyChatService_BroadcastAllMessagesServer) error {
	log.Printf(req.User.Name + " broadcasts all messages")
	srLock.Lock()
	check, exists := server.currentUsers[req.User.UserID]

	//Checks wether or not the user already exists, if it does, update the stream value
	if exists {
		check.broadCastAll = stream
		server.currentUsers[req.User.UserID] = check
	} else {
		server.currentUsers[req.User.UserID] = streamsForClients{
			broadCastAll: stream}
	}
	srLock.Unlock()

	//sets up a timer, that executes at a certain interval
	//timer := time.NewTicker(4 * time.Second)

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
		//may need to go back to different loop format: _,user := range s.currentUsers
		err := server.currentUsers[i].broadCastAll.Send(&proto.BroadcastAllResponse{
			Messages:  tmpMessages,
			TimeStamp: int32(server.lambortTime),
		})

		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (server *ChittyChatServiceServer) BroadcastJoin(req *proto.NewClientJoinedRequest, stream proto.ChittyChatService_BroadcastJoinServer) error {
	log.Printf(req.User.Name + " joins the chat")

	//Server updates the users ID. Currently it does not take into account if one has left the
	srLock.Lock()
	if totalAmuntUsers == 0 {
		totalAmuntUsers += 1
		req.User.UserID = 0
	} else {
		req.User.UserID = int32(totalAmuntUsers)
		totalAmuntUsers += 1
	}
	srLock.Unlock()

	srLock.Lock()
	check, exists := server.currentUsers[req.User.UserID]
	if exists {
		//if the user already exists, we update the stream
		check.broadCastJoin = stream
		server.currentUsers[req.User.UserID] = check
	} else {
		server.currentUsers[req.User.UserID] = streamsForClients{
			broadCastJoin: stream}
	}
	srLock.Unlock()

	var message = "User " + req.User.Name + " joined at Lamport Time " + strconv.Itoa(server.lambortTime)

	//sends the join message to all current users,
	for i := range server.currentUsers {
		//may need to go back to different loop format: _,user := range s.currentUsers
		err := server.currentUsers[i].broadCastJoin.Send(&proto.NewClientJoinedResponse{
			Message:   message,
			TimeStamp: int32(server.lambortTime),
		})
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}

func (server *ChittyChatServiceServer) BroadcastLeave(req *proto.ClientLeaveRequest, stream proto.ChittyChatService_BroadcastLeaveServer) error {
	log.Printf(req.User.Name + " leaves")
	srLock.Lock()
	check, exists := server.currentUsers[req.User.UserID]
	if exists {
		//if the user already exists, we update the stream
		check.broadCastLeave = stream
		server.currentUsers[req.User.UserID] = check
	} else {
		server.currentUsers[req.User.UserID] = streamsForClients{
			broadCastLeave: stream}
	}
	srLock.Unlock()

	var message = "User " + req.User.Name + " left at Lamport Time " + strconv.Itoa(server.lambortTime)

	//sends the leave message to all current users,
	for i := range server.currentUsers {
		//may need to go back to different loop format: _,user := range s.currentUsers
		err := server.currentUsers[i].broadCastLeave.Send(&proto.ClientLeaveResponse{
			Message:   message,
			TimeStamp: int32(server.lambortTime),
		})
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	//removes client from current clients
	delete(server.currentUsers, req.User.UserID)

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
