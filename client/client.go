package main

import (
	"context"
	proto "handin3/grpc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	/*
		conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Not working")
		}

		client := proto.NewChittyChatServiceClient(conn)

		if err != nil {
			log.Fatalf("Not working")
		}

		var pls, _ = client.PublishMessage(context.Background(), &proto.PublishRequest{
			User:      "miguel",
			Message:   "WOOOOP",
			TimeStamp: 12,
		})

		fmt.Println(pls.Success, pls.Message)

		//go client()
	*/

}

func client() {
	var lamportTime int = 0

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := proto.NewChittyChatServiceClient(conn)

	if err != nil {
		log.Fatalf("Not working")
	}

	lamportTime += 1
	client.PublishMessage(context.Background(), &proto.PublishRequest{
		User:      "miguel",
		Message:   "Hello",
		TimeStamp: int32(lamportTime),
	})

}
