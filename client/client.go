package main

import (
	"context"
	"fmt"
	proto "handin3/grpc"
	"io"
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
//3) done/doing
//4) to do (might just be to add a logPrint everytime something happens)
//5) to do
//6) done/doing (this can be done, but not like the assignments wants)
//7) done/doing (this can be done, but not like the assignments wants)

var theLog = []proto.PostMessage{}

func main() {
	var lamportTime int = 0

	conn, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	client := proto.NewChittyChatServiceClient(conn)

	if err != nil {
		log.Fatalf("Not working")
	}

	//THE LOG WHICH IS GONNA BE "LOGGED" TO
	//var Log    []proto.PublishRequest

	//lamportTime += 1

	user := proto.User{
		Name:   "Morten",
		UserID: 1,
	}

	//fmt.Println(faf.Message, " ", faf.TimeStamp)

	//lamportTime = compareLT(lamportTime, int(work.TimeStamp))
	//lamportTime += 1

	//fmt.Println(work.Message, " ", work.TimeStamp)

	lamportTime += 1
	var pls, _ = client.PublishMessage(context.Background(), &proto.PostMessage{
		User:      &user,
		Message:   "Hej",
		TimeStamp: int32(lamportTime),
	})

	fmt.Println(pls.Success, " ", pls.Message, " ")

	lamportTime = compareLT(lamportTime, int(pls.TimeStamp))
	lamportTime += 1

	fmt.Println("Client Lamport Time ", lamportTime)

	//testing if the server can handle a message that is too long
	/*
		lamportTime += 1
		var pls2, _ = client.PublishMessage(context.Background(), &proto.PostMessage{
			User:      &user,
			Message:   "HEJ ALLESAMMEN MIT NAVN ER LUKAS OG JEG ER EN Dasdadasasasasasasasasasasasasasasasasasasasasasasasasasasasdadasdwrfqwgqqgqqqwsadsadwawqfqqweqweqwsdadwadafdqwfwqqfqfqfqfafsafwasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasasas",
			TimeStamp: int32(lamportTime),
		})


		fmt.Println(pls2.Success, " ", pls2.Message, " ")

		lamportTime = compareLT(lamportTime, int(pls2.TimeStamp))
		lamportTime += 1

		fmt.Println("Client Lamport Time ", lamportTime)
	*/

	//NÅR VI ER NÅET TIL AT LAVE BROADCAST LEAVE UDKOMMENTER OG ARBEJD VIDERE!!!
	/*
		lamportTime += 1
		var gooo, _ = client.BroadcastLeave(context.Background(), &proto.ClientLeaveRequest{
			User:      &user,
			TimeStamp: int32(lamportTime),
		})

		fmt.Println(gooo.Message, " ", gooo.TimeStamp)

		lamportTime = compareLT(lamportTime, int(gooo.TimeStamp))
		lamportTime += 1

		fmt.Println("Client Lamport Time ", lamportTime)
	*/

	//Stream starts here
	var stream, erro = client.BroadcastAllMessages(context.Background(), &proto.BroadcastAllRequest{})

	var broadcastJoin, _ = client.BroadcastJoin(context.Background(), &proto.NewClientJoinedRequest{
		User:      &user,
		TimeStamp: int32(lamportTime),
	})

	//Creating a timer, to run the stream every 5 seconds
	stop := time.NewTicker(7 * time.Second)

	if erro != nil {
		log.Fatalf("Not working")
	}

	//Setting up client reading stream
	lamportTime += 1

	running := true

	for running {
		select {
		case <-stop.C:
			err := stream.CloseSend()
			if err != nil {
				log.Fatalf("Failed to close stream: %v", err)
			}
			running = false

		default:
			getMessage, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					running = false
					break
				}
				log.Fatalf("Failed to receive a message: %v", err)
			}
			// Should be added to the log
			//fmt.Println(getMessage.Messages.User.Name, ", ", getMessage.Messages.Message, ", ", getMessage.Messages.TimeStamp)

			// Checking if the message is already in the log
			var messageInLog = false

			for i := range theLog {
				if theLog[i].User.UserID == getMessage.Messages.User.UserID && theLog[i].Message == getMessage.Messages.Message && theLog[i].TimeStamp == getMessage.Messages.TimeStamp {
					messageInLog = true
					break
				}
			}

			if !messageInLog {
				theLog = append(theLog, proto.PostMessage{
					User:      getMessage.Messages.User,
					Message:   getMessage.Messages.Message,
					TimeStamp: int32(lamportTime),
				})
			}

			getJoin, err := broadcastJoin.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Failed to receive join message: %v", err)
			}
			fmt.Println(getJoin.Message, " ", getJoin.TimeStamp)

			broadcastJoin.CloseSend()
		}

		lamportTime = compareLT(lamportTime, int(pls.TimeStamp))
		lamportTime += 1
	}

	fmt.Println("now for the log")
	for i := range theLog {
		message := &theLog[i]
		fmt.Println(message.User.Name, ", ", message.Message, ", ", message.TimeStamp)
	}
}

func compareLT(thisLT int, otherLT int) int {
	var result int

	if thisLT > otherLT {
		result = thisLT
	} else {
		result = otherLT
	}

	return result
}
