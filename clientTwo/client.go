package main

import (
	"bufio"
	"context"
	"fmt"
	proto "handin3/grpc"
	"log"
	"os"
	"sync"
	"unicode/utf8"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var theLog = []proto.PostMessage{}

var lamportTime int = 0

var thisUser proto.User

var BroadcastStream proto.ChittyChatService_ConnectedClient

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

	//DER SKAL LAVES NOGET HVOR CLIENTEN KAN MODTAGE BROADCASTS

	for {
		if !loggedIn {
			fmt.Println("Enter your name: ")
			var name string

			fmt.Scanln(&name)

			thisUser = proto.User{
				Name:   name,
				UserID: 1,
			}

			joinResponse, err := client.JoinServer(context.Background(), &proto.JoinRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			if err != nil {
				log.Fatalf("Failed to join server: %v", err)
			}

			//updates the ID of the user
			thisUser.UserID = joinResponse.UserID

			log.Println("Logging in as ", name)
			var message = thisUser.Name + " joined at Lamport Time: " + string(lamportTime)
			BroadcastStream, _ = client.Connected(context.Background())
			BroadcastStream.Send(&proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})

			go GetMessages()
			loggedIn = true

		}

		var input string

		fmt.Println("Hi", thisUser.Name, ", please enter what you want to do ")
		fmt.Println("type 'help' to list all commands")

		fmt.Scanln(&input)

		if input == "help" {
			fmt.Println("type 'send' to send a message")
			fmt.Println("type 'quit' to quit")
			fmt.Println("type 'profile' to see your profile")
			fmt.Println("type 'log' to see the log")
			continue
		}

		if input == "send" {
			fmt.Println("Enter your message: ")
			message, _ := reader.ReadString('\n')

			if !checkMessage(message) {
				//If the message is invalid, we skip the rest of the loop and wait new input
				continue
			} else {
				log.Printf("Sending message...")
				BroadcastStream.Send(&proto.PostMessage{
					User:      &thisUser,
					Message:   message,
					TimeStamp: int32(lamportTime),
				})

			}
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

			BroadcastStream.Send(&proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})

			log.Printf(thisUser.Name + " logging out")
			break
		}
	}
}

func GetMessages() {
	for {
		GetMessages, _ := BroadcastStream.Recv()
		if GetMessages != nil {
			theLog = append(theLog, *GetMessages)
			log.Printf("received broadcast messages")
		}
	}
}

func compareLamportTime(lamportTime int, messageLamportTime int) int {
	if lamportTime > messageLamportTime {
		return lamportTime + 1
	} else {
		lamportTime = messageLamportTime
		return lamportTime + 1
	}
}

func checkMessage(message string) bool {
	if !utf8.ValidString(message) {
		log.Printf("Messages must be valid in utf8")
		return false
	} else if len(message) > 128 {
		log.Printf("Messages cant be longer than 128 characters")
		return false

	} else if len(message) == 0 {
		log.Printf("Message must not be empty")
		return false

	} else {
		return true
	}
}
