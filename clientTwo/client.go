package main

import (
	"bufio"
	"context"
	"fmt"
	proto "handin3/grpc"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var theLog = []proto.PostMessage{}

var lamportTime int = 0

var thisUser proto.User

var broadcastJoin proto.ChittyChatService_BroadcastJoinClient

var broadcastLeave proto.ChittyChatService_BroadcastLeaveClient

var BroadcastAllMessages proto.ChittyChatService_BroadcastAllMessagesClient

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
			broadcastJoin, _ = client.BroadcastJoin(context.Background(), &proto.NewClientJoinedRequest{
				User:      &thisUser,
				TimeStamp: 0,
			})

			loggedIn = true
		}

		var input string

		fmt.Println("Hi", thisUser.Name, ", please enter what you want to do ")
		fmt.Println("type 'list' to list all options")

		fmt.Scanln(&input)

		if input == "list" {
			fmt.Println("type 'send' to send a message")
			fmt.Println("type 'quit' to quit")
			fmt.Println("type 'profile' to see your profile")
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
			fmt.Println("Profile: ")
			fmt.Println("Name: ", thisUser.Name)
			fmt.Println("UserID: ", thisUser.UserID)
		}

		if input == "quit" {
			broadcastLeave, _ = client.BroadcastLeave(context.Background(), &proto.ClientLeaveRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			//Her skal den håndtere broadcasten over at den selv forlader serveren

			log.Printf("Logging out")
			break
		}
	}
}
