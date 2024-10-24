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

var lamportTime int = 0

var thisUser proto.User

var BroadcastStream proto.ChittyChatService_CommunicateClient

var clLock sync.Mutex
var stop bool

//NOTE: We need to use goroutines to handle the broadcasts
//NOTE: We may need to make more rpc methods, for out clients joining and leaving (splitting the broadcast into two methods), as we did before
//NOTE: Still missing lamportTime.

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	reader := bufio.NewScanner(os.Stdin)
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

			lamportTime++

			fmt.Scanln(&name)
			/*for reader.Scan() {
				name = reader.Text()
				break
			}*/

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
				lamportTime++
			} else {
				log.Println(joinResponse.Message)
				lamportTime++
			}

			thisUser.UserID = joinResponse.UserID

			lamportTime++
			BroadcastStream, _ = client.Communicate(context.Background())

			go GetMessages()
			loggedIn = true
			fmt.Println("Hi", thisUser.Name, ", Type anything to write a message")
			fmt.Println("Here is a list of commands")
			fmt.Println("Type '/profile' to see your profile")
			fmt.Println("type '/log' to see the log")
			fmt.Println("Type '/quit' to quit")
			fmt.Println("Type '/help' for list of commands")
		}

		var input string

		for reader.Scan() {
			input = reader.Text()
			break
		}

		if input == "/help" {
			lamportTime++
			fmt.Println("Here is a list of commands")
			fmt.Println("Type '/profile' to see your profile")
			fmt.Println("type '/log' to see the log")
			fmt.Println("Type '/quit' to quit")
			fmt.Println("Type '/help' for list of commands")
			continue
		} else if input == "/profile" {
			lamportTime++
			fmt.Println("Your Profile: ")
			fmt.Println("Name: ", thisUser.Name)
			fmt.Println("UserID: ", thisUser.UserID)
		} else if input == "/quit" {
			lamportTime++
			LeaveResponse, err := client.LeaveServer(context.Background(), &proto.LeaveRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			if err != nil {
				log.Fatalf("Failed to leave server: %v", err)
				lamportTime++
			}

			log.Printf(LeaveResponse.Message)
			break
		} else {
			if !checkMessage(input) {
				//If the message is invalid, we skip the rest of the loop and wait new input
				continue
			} else {
				lamportTime++
				actualMessage := thisUser.Name + ": " + input
				log.Printf("Sending message...")
				BroadcastStream.Send(&proto.PostMessage{
					User:      &thisUser,
					Message:   actualMessage,
					TimeStamp: int32(lamportTime),
				})

			}
		}
	}
}

func GetMessages() {
	var stop bool = true
	for stop {
		GetMessages, err := BroadcastStream.Recv()
		if err != nil {
			//log.Printf("Failed to receive broadcast messages: %v", err)
			lamportTime++
			stop = false
			log.Fatal(err)
		} else {
			clLock.Lock()
			log.Printf(GetMessages.Message)
			lamportTime++
			clLock.Unlock()
			//log.Printf("received broadcast messages")
		}
	}
}

func checkMessage(message string) bool {
	if !utf8.ValidString(message) {
		log.Printf("Messages must be valid in utf8")
		lamportTime++
		return false
	} else if len(message) > 128 {
		log.Printf("Messages cant be longer than 128 characters")
		lamportTime++
		return false

	} else if len(message) == 0 {
		log.Printf("Message must not be empty")
		lamportTime++
		return false

	} else {
		return true
	}
}
