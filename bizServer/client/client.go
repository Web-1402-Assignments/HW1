package main

import (
	pb "bizServer"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:5062", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new UserService client
	client := pb.NewUserServiceClient(conn)

	// Create a request for the first service: get_users
	getUsersReq := &pb.GetUsersRequest{
		UserId:    123,
		AuthKey:   "valid_key",
		MessageId: 2,
	}

	// Send the request to the server
	getUsersRes, err := client.GetUsers(context.Background(), getUsersReq)
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}

	// Process the response
	for _, user := range getUsersRes.Users {
		log.Printf("User ID: %d, Name: %s", user.Id, user.Name)
	}

	log.Printf("Message ID: %d", getUsersRes.MessageId)

	// Create a request for the second service: get_users_with_sql_inject
	getUsersWithSQLInjectReq := &pb.GetUsersWithSQLInjectRequest{
		UserId:    "1 OR 1=1; --",
		AuthKey:   "valid_key",
		MessageId: 4,
	}

	// Send the request to the server
	getUsersWithSQLInjectRes, err := client.GetUsersWithSQLInject(context.Background(), getUsersWithSQLInjectReq)
	if err != nil {
		log.Fatalf("Failed to get users with SQL injection: %v", err)
	}

	// Process the response
	for _, user := range getUsersWithSQLInjectRes.Users {
		log.Printf("User ID: %d, Name: %s", user.Id, user.Name)
	}

	log.Printf("Message ID: %d", getUsersWithSQLInjectRes.MessageId)
}
