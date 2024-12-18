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

	for {
		if !loggedIn {
			fmt.Println("Enter your name: ")
			var name string

			//Name counts as event
			clLock.Lock()
			lamportTime++
			clLock.Unlock()

			for reader.Scan() {
				name = reader.Text()
				break
			}
			thisUser = proto.User{
				Name:   name,
				UserID: 1,
			}

			//Joining the server
			clLock.Lock()
			lamportTime++
			clLock.Unlock()
			log.Printf("Attemting to connect to server at Lamport time: %d", lamportTime)
			joinResponse, err := client.JoinServer(context.Background(), &proto.JoinRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			if err != nil {
				log.Fatalf("Failed to join server: %v", err)
				clLock.Lock()
				lamportTime++
				clLock.Unlock()
			} else {
				//Getting response from server
				lamportTime = compareLamportTime(int(joinResponse.TimeStamp))
				lamportTime++
				log.Println(joinResponse.Message, lamportTime)
				thisUser.UserID = joinResponse.UserID
			}

			clLock.Lock()
			lamportTime++

			log.Printf("Starting to communicate with server... %d", lamportTime)
			BroadcastStream, _ = client.Communicate(context.Background())
			BroadcastStream.Send(&proto.PostRequest{
				User:      &thisUser,
				Message:   "Connected to server",
				TimeStamp: int32(lamportTime),
			})

			clLock.Unlock()

			go GetMessages()
			loggedIn = true
			fmt.Println("Hi", thisUser.Name, ", Type anything to write a message")
			fmt.Println("Here is a list of commands")
			fmt.Println("Type '/profile' to see your profile")
			fmt.Println("Type '/lt' to see the current lamport time")
			fmt.Println("Type '/quit' to quit")
			fmt.Println("Type '/help' for list of commands")
		}

		var input string

		for reader.Scan() {
			input = reader.Text()
			break
		}

		if input == "/help" {
			clLock.Lock()
			lamportTime++
			clLock.Unlock()

			fmt.Println("Here is a list of commands")
			fmt.Println("Type '/profile' to see your profile")
			fmt.Println("type '/log' to see the log")
			fmt.Println("Type '/lt' to see the current lamport time")
			fmt.Println("Type '/quit' to quit")
			fmt.Println("Type '/help' for list of commands")
			continue
		} else if input == "/profile" {
			clLock.Lock()
			lamportTime++
			clLock.Unlock()

			fmt.Println("Your Profile: ")
			fmt.Println("Name: ", thisUser.Name)
			fmt.Println("UserID: ", thisUser.UserID)
		} else if input == "/quit" {
			clLock.Lock()
			lamportTime++
			clLock.Unlock()

			log.Printf("Leaving server... at: %d", lamportTime)
			LeaveResponse, err := client.LeaveServer(context.Background(), &proto.LeaveRequest{
				User:      &thisUser,
				TimeStamp: int32(lamportTime),
			})

			clLock.Lock()
			lamportTime = compareLamportTime(int(LeaveResponse.TimeStamp))
			lamportTime++
			clLock.Unlock()

			if err != nil {
				log.Fatalf("Failed to leave server: %v", err)
				clLock.Lock()
				lamportTime++
				clLock.Unlock()
			}

			log.Printf(LeaveResponse.Message)
			log.Printf("Left server at Lamport Time: %d", lamportTime)

			break
		} else if input == "/lt" {
			log.Printf("Lamport Time: %v", lamportTime)
		} else {
			if !checkMessage(input) {
				//If the message is invalid, we skip the rest of the loop and wait new input
				continue
			} else {
				lamportTime++
				actualMessage := thisUser.Name + " says: " + input
				log.Printf("Sending message at Lamport Time: %d", lamportTime)
				BroadcastStream.Send(&proto.PostRequest{
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
		//Recieve message from server
		GetMessages, err := BroadcastStream.Recv()
		if err != nil {
			log.Printf("Server has shut down: %v", err)
			lamportTime++
			stop = false
			log.Fatal(err)
		}

		if GetMessages != nil {
			clLock.Lock()
			lamportTime = compareLamportTime(int(GetMessages.TimeStamp))
			lamportTime++
			clLock.Unlock()
			log.Printf(GetMessages.Message + fmt.Sprint(lamportTime))
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
		log.Printf("Message cant be empty")
		lamportTime++
		return false
	} else {
		return true
	}
}

func compareLamportTime(otherLamportTime int) int {
	if lamportTime > otherLamportTime {
		return lamportTime
	} else if otherLamportTime > lamportTime {
		return otherLamportTime
	} else {
		return lamportTime
	}
}
