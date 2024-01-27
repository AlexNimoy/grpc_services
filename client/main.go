package main

import (
	pb "client/proto"
	"context"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	for {
		// Generate and send a random message
		message := generateRandomMessage()
		response, err := client.SendMessage(context.Background(), &pb.Message{Text: message})
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		log.Printf("Server Response: %s", response.Reply)

		time.Sleep(3 * time.Second)
	}
}

func generateRandomMessage() string {
	messages := []string{"Hello", "How are you?", "Sample message", "Testing gRPC"}
	return messages[rand.Intn(len(messages))]
}
