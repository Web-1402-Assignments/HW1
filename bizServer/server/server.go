package main

import (
	pb "bizServer"
	"errors"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"database/sql"
)

type userServiceServer struct {
	// Implement the necessary fields for your server
	pb.UnimplementedUserServiceServer // Embed the UnimplementedUserServiceServer type
	db *sql.DB
}

// Implement the first service: get_users
func (s *userServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	/////////////////////////////////////// getter
	if req.AuthKey != "valid_key" {
		return nil, errors.New("invalid auth_key")
	}

	if req.MessageId%2 != 0 || req.MessageId <= 0 {
		return nil, errors.New("invalid message_id")
	}

	// TODO: Implement the database query to fetch users based on the provided input parameters
	// Use s.db to execute the query

	// Example response
	users := []*pb.User{
		{Id: 1, Name: "User 1"},
		{Id: 2, Name: "User 2"},
		// Add more users as needed
	}

	return &pb.GetUsersResponse{
		Users:     users,
		MessageId: req.MessageId + 1,
	}, nil
}

// Implement the second service: get_users_with_sql_inject
func (s *userServiceServer) GetUsersWithSQLInject(ctx context.Context, req *pb.GetUsersWithSQLInjectRequest) (*pb.GetUsersWithSQLInjectResponse, error) {
	if req.AuthKey != "valid_key" {
		return nil, errors.New("invalid auth_key")
	}

	if req.MessageId%2 != 0 || req.MessageId <= 0 {
		return nil, errors.New("invalid message_id")
	}

	// TODO: Implement the database query to fetch users based on the provided input parameters
	// IMPORTANT: Ensure you properly sanitize the input to prevent SQL injection vulnerabilities
	// Use s.db to execute the query

	// Example response
	users := []*pb.User{
		{Id: 1, Name: "User 1"},
		{Id: 2, Name: "User 2"},
		// Add more users as needed
	}

	return &pb.GetUsersWithSQLInjectResponse{
		Users:     users,
		MessageId: req.MessageId + 1,
	}, nil
}

func main() {
	// TODO: Initialize your Postgres database connection
	// Replace the connection string with your actual Postgres connection details
	
	// db, err := sql.Open("postgres", "host=127.0.0.1 port=5062 user=your_user password=your_password dbname=your_db sslmode=disable")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v", err)
	// }
	// defer db.Close()

	// TODO: Ensure the necessary tables are created in the database

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the UserService server
	// userService := &userServiceServer{db: db}
	userService := &userServiceServer{}
	pb.RegisterUserServiceServer(server, userService) // Register the UserService implementation

	// Start listening on a TCP port
	listener, err := net.Listen("tcp", ":5062")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the gRPC server
	log.Println("Server started, listening on port 5062")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}




// var (
// 	port = flag.Int("port", 5062, "The server port")
// )

// func (s *server) GetUser(ctx context.Context, request *pb.Get) (*pb.UserResponse, error) {
// 	// Handle the GetUser request and return a response or an error
// 	// Your implementation goes here
// }

// func main() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGetUsersServer(s, &server{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// 	/*flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	var opts []grpc.ServerOption

// 	grpcServer := grpc.NewServer(opts...)
// 	pb.RegisterGetUsersServer(grpcServer, newServer())
// 	grpcServer.Serve(lis)*/
// }
