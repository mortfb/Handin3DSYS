package main

import (
	"context"
	"fmt"
	proto "handin3/grpc"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//TODO: Make it so the client is "manual", sp it takes manual input from client, instead of prewritten
//		Make streams (this is porbably what we should do) so server can broadcast Messages from server to Clients, this could maybe be done with a pointer to the Message[] in the server
//		How can a client broadcast, when expecting input.
//		How does a client know that someone else posted before it finished
//		The client needs to update its own lamportTime if the server has a greate lamport time (like it is done in server)

//CHECKLIST:
//R1 in progress
//R2 should be done
//R3 to do
//R4 to do
//R5 Technically done, but need more functionality
//R6 in progress, but need to figure out how to broadcast first
//R7 to do
//R8 in progress, but need to figure out how to broadcast first

//TECHNICAL REQUIERMENTS
//1) done/doing
//2) done/doing
//3) to do
//4) to do (might just be to add a logPrint everytime something happens)
//5) to do
//6) done/doing (this can be done, but not like the assignments wants)
//7) done/doing (this can be done, but not like the assignments wants)

func main() {
	var lamportTime int = 0

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := proto.NewChittyChatServiceClient(conn)

	if err != nil {
		log.Fatalf("Not working")
	}

	//THE LOG WHICH IS GONNA BE "LOGGED" TO
	//var Log    []proto.PublishRequest

	lamportTime += 1

	user := proto.User{
		Name:   "Lukas",
		UserID: 1,
	}

	var work, _ = client.NewClientJoined(context.Background(), &proto.NewClientJoinedRequest{
		User:      &user,
		TimeStamp: int32(lamportTime),
	})

	fmt.Println(work.Message, " ", work.TimeStamp)

	lamportTime += 1
	var pls, _ = client.PublishMessage(context.Background(), &proto.PostMessage{
		User:      &user,
		Message:   "WOOOOP",
		TimeStamp: int32(lamportTime),
	})

	fmt.Println(pls.Success, " ", pls.Message, " ")

	fmt.Println("Client Lamport Time ", lamportTime)

	lamportTime += 1
	var gooo, _ = client.ClientLeave(context.Background(), &proto.ClientLeaveRequest{
		User:      &user,
		TimeStamp: int32(lamportTime),
	})

	fmt.Println(gooo.Message, " ", gooo.TimeStamp)

	fmt.Println("Client Lamport Time ", lamportTime)

	//Stream starts here
	var stream, erro = client.BroadcastAllMessages(context.Background(), &proto.BroadcastAllRequest{})

	//Creating a timer, to run the stream every 5 seconds
	stop := time.NewTicker(5 * time.Second)

	if erro != nil {
		log.Fatalf("Not working")
	}

	//Setting up client reading stream
	for {
		select {
		case <-stop.C:
			err := stream.CloseSend()
			if err != nil {
				log.Fatalf("Failed to close stream: %v", err)
			}
			return

		default:
			message, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive a message : %v", err)
			}
			//Should be added to the log
			fmt.Println(message.Messages.User.Name, ", ", message.Messages.Message, ", ", message.Messages.TimeStamp)

		}

	}

	for {
		message, _ := stream.Recv()

		fmt.Println(message.Messages.User.Name, " ", message.Messages.Message)
	}

}
