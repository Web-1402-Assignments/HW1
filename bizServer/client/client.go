package main

import (
	pb "HW1"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5062", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Create a request to get users
	getUsersReq := &pb.GetUsersRequest{
		UserId:    1111,
		AuthKey:   "valid_key",
		MessageId: 2,
	}

	getUsersRes, err := client.GetUsers(context.Background(), getUsersReq)
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}

	for _, user := range getUsersRes.Users {
		log.Printf("User ID: %d,	Name: %s ,	Family: %s ,	Age: %d,	Sex: %s ,	Time: %s", 
				user.Id, user.Name, user.Family, user.Age, user.Sex, user.CreatedAt)
	}
	log.Printf("Message ID: %d", getUsersRes.MessageId)

	// Create a request with sql inject
	getUsersWithSQLInjectReq := &pb.GetUsersWithSQLInjectRequest{
		UserId:    "1111 OR 1=1 --",
		AuthKey:   "valid_key",
		MessageId: 8,
	}

	// Send the request to the server
	getUsersWithSQLInjectRes, err := client.GetUsersWithSQLInject(context.Background(), getUsersWithSQLInjectReq)
	if err != nil {
		log.Fatalf("Failed to get users with SQL injection: %v", err)
	}

	for _, user := range getUsersWithSQLInjectRes.Users {
		log.Printf("User ID: %d,	Name: %s ,	Family: %s ,	Age: %d,	Sex: %s ,	Time: %s", 
				user.Id, user.Name, user.Family, user.Age, user.Sex, user.CreatedAt)
	}

	log.Printf("Message ID: %d", getUsersWithSQLInjectRes.MessageId)
}
