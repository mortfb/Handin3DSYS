package main

import (
	"context"
	proto "handin3/grpc"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChittyChatServiceServer struct {
	proto.UnimplementedChittyChatServiceServer

	currentUsers []proto.ChittyChatService_ConnectedServer

	totalAmountUsers int32

	lamportTime int

	messages []proto.PostMessage
}

var srLock sync.Mutex
var waitGroup sync.WaitGroup

func main() {
	server := &ChittyChatServiceServer{
		currentUsers:     []proto.ChittyChatService_ConnectedServer{},
		totalAmountUsers: 0,
		lamportTime:      0,
		messages:         []proto.PostMessage{},
	}

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

func (server *ChittyChatServiceServer) JoinServer(ctx context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
	srLock.Lock()
	server.lamportTime = compareLamportTime(server.lamportTime, int(req.TimeStamp))

	var response = &proto.JoinResponse{
		UserID:    server.totalAmountUsers,
		TimeStamp: int32(server.lamportTime),
	}
	server.totalAmountUsers++
	srLock.Unlock()
	return response, nil
}

/*func sendMessages(stream proto.ChittyChatService_ConnectedServer, server *ChittyChatServiceServer) {
	for {
		message, err := stream.Recv()
		if err == nil {
			log.Printf(message.User.Name + "%s: %s" + message.Message)
			for i := range server.currentUsers {
				if i != int(message.User.UserID) { //userID
					server.currentUsers[i].Send(&proto.PostResponse{
						User:      message.User,
						Message:   message.Message,
						TimeStamp: message.GetTimeStamp(),
					})
				}
			}
		} else {
			//log.Println("Did not receive any messages")
		}

	}
}*/

func (server *ChittyChatServiceServer) Connected(stream proto.ChittyChatService_ConnectedServer) error {
	//Lige nu bruges req ikke til noget, hvad skal den bruges til?
	server.currentUsers = append(server.currentUsers, stream)
	//Hvorfor gøre dette?

	for {
		message, err := stream.Recv()
		if err == nil {
			log.Printf(message.Message)
			for i := range server.currentUsers {
				if i != int(message.User.UserID) { //userID
					server.currentUsers[i].Send(&proto.PostResponse{
						User:      message.User,
						Message:   message.Message,
						TimeStamp: message.GetTimeStamp(),
					})
				}
			}
		} else {
			//log.Println("Did not receive any messages")
		}

	}

	//Hivs man kan gøre dette i stedet?
	//Ved ikke om det her er bedre/være end at bruge goroutinen
	/*for {
		message, err := stream.Recv()
		if err == nil {
			srLock.Lock()
			server.lamportTime = compareLamportTime(server.lamportTime, int(message.TimeStamp))
			server.messages = append(server.messages, *message)
			srLock.Unlock()

			//Var server.sendMessages(stream) før, men den brokkede sig over at den ikke blev brugt ovre i client
			//DET KAN GODT VÆRE DEN SKAL ÆNDRES TILBAGE, HVIS DET VIRKER BEDRE

			//Den skal vel kun bruges her i server, derfor kan skal man ikke have at den kaldes på en (server *ChittyChatServiceServer), men bare som "en normal" metode
			//Har også ændret det i selve metoden, så den bare tager den server, som den skal bruge
			sendMessages(stream, server)
		}
	}*/
	return nil
}

func compareLamportTime(lamportTime int, messageLamportTime int) int {
	if lamportTime > messageLamportTime {
		return lamportTime + 1
	} else {
		lamportTime = messageLamportTime
		return lamportTime + 1
	}
}
