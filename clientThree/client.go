package main

import (
	proto "Handin3DSYS/grpc"
	"bufio"
	"context"
	"fmt"
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

//NOTE: We need to use goroutines to handle the broadcasts
//NOTE: We may need to make more rpc methods, for out clients joining and leaving (splitting the broadcast into two methods), as we did before
//NOTE: Still missing lamportTime.

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

			log.Println("Attempting to join server...")
			joinResponse, err := client.JoinServer(context.Background(), &proto.JoinRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})
			if err != nil {
				log.Fatalf("Failed to join server: %v", err)
			}
			jointmp, err := joinResponse.Recv()

			if err != nil {
				log.Fatalf("Failed to join server: %v", err)
			} else {
				log.Println(jointmp.Message)
			}

			thisUser.UserID = jointmp.UserID
			//var message = thisUser.Name + " joined at Lamport Time: " + string(lamportTime)
			BroadcastStream, _ = client.Connected(context.Background())
			/*BroadcastStream.Send(&proto.PostMessage{
				User:      &thisUser,
				Message:   message,
				TimeStamp: int32(lamportTime),
			})
			*/
			go GetMessages()
			loggedIn = true

		}

		var input string

		fmt.Println("")
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
			if len(theLog) == 0 {
				fmt.Println("No messages in log")
			} else {
				clLock.Lock()
				for i := range theLog {
					fmt.Println(theLog[i].User.Name, ": ", theLog[i].Message, ", ", theLog[i].TimeStamp)
				}
				clLock.Unlock()
			}
		}

		if input == "quit" {
			//Her skal den håndtere broadcasten over at den selv forlader serveren
			LeaveResponse, err := client.LeaveServer(context.Background(), &proto.LeaveRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			if err != nil {
				log.Fatalf("Failed to leave server: %v", err)
			}

			leaveMessage, _ := LeaveResponse.Recv()

			log.Printf(leaveMessage.Message)
			break
		}

	}
}

func GetMessages() {
	for {
		GetMessages, err := BroadcastStream.Recv()
		if err != nil {
			log.Printf("Failed to receive broadcast messages: %v", err)
		} else {
			clLock.Lock()
			fmt.Println(GetMessages.Message, ", ", GetMessages.TimeStamp)
			clLock.Unlock()
			log.Printf("received broadcast messages")
		}
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
