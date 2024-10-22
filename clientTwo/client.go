package main

import (
	"bufio"
	"context"
	"fmt"
	proto "handin3/grpc"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var theLog = []proto.PostMessage{}

var lamportTime int = 0

var thisUser proto.User

var BroadcastStream proto.ChittyChatService_BroadcastMessagesClient

var clLock sync.Mutex

// Channels we need to use, so the client knows when to open and receive broadcasts, so our program dont crash. These needs to be used in a select statement
var broadCastAllChan = make(chan []*proto.PostMessage)

//NOTE: We need to use goroutines to handle the broadcasts
//NOTE: We may need to make more rpc methods, for out clients joining and leaving (splitting the broadcast into two methods), as we did before
//NOTE: Still missing lamportTime.

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	reader := bufio.NewReader(os.Stdin)
	client := proto.NewChittyChatServiceClient(conn)

	if err != nil {
		log.Fatalf("Not working")
	}

	var loggedIn bool = false

	for {
		if !loggedIn {
			fmt.Println("Enter your name: ")
			var name string

			fmt.Scanln(&name)

			thisUser = proto.User{
				Name:   name,
				UserID: 1,
			}

			log.Println("Logging in as ", name)
			BroadcastStream, _ = client.BroadcastMessages(context.Background(), &proto.BroadcastRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			var message = thisUser.Name + " joined at Lamport Time: "
			client.PublishMessage(context.Background(), &proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})

			loggedIn = true
		}

		var input string
		GetMessages, _ := BroadcastStream.Recv()
		var gofor []*proto.PostMessage
		if GetMessages != nil {
			var gofor = GetMessages.Messages
			log.Printf("received broadcast messages")
			if gofor == nil {
				log.Printf("recieved messages are null")
			}
		}
		for i := range gofor {
			fmt.Printf("Message: %s", gofor[i].Message)
		}

		fmt.Println("Hi", thisUser.Name, ", please enter what you want to do ")
		fmt.Println("type 'list' to list all commands")

		fmt.Scanln(&input)

		if input == "list" {
			fmt.Println("type 'send' to send a message")
			fmt.Println("type 'quit' to quit")
			fmt.Println("type 'profile' to see your profile")
			fmt.Println("type 'log' to see the log")
			continue
		}

		if input == "send" {
			fmt.Println("Enter your message: ")
			message, _ := reader.ReadString('\n')

			var msg, _ = client.PublishMessage(context.Background(), &proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})

			if !msg.Success {
				fmt.Println(msg.Message)
			}

			//Brug msg til at opdatere lamportTime
		}

		if input == "profile" {
			fmt.Println("Your Profile: ")
			fmt.Println("Name: ", thisUser.Name)
			fmt.Println("UserID: ", thisUser.UserID)
		}

		if input == "log" {
			for i := range theLog {
				message := &theLog[i]
				fmt.Println(message.User.Name, ": ", message.Message, ", ", message.TimeStamp)
			}
		}

		if input == "quit" {
			//Her skal den håndtere broadcasten over at den selv forlader serveren
			var message = thisUser.Name + " left at Lamport Time: "
			client.PublishMessage(context.Background(), &proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})

			log.Printf("Logging out")
			break
		}

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
