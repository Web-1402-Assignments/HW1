package main

import (
	pb "HW1/proto"
	"context"
	"errors"
	"log"
	"net"
	"time"
	"regexp"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer 
	db *sql.DB
}
var (
	client = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
	})
	redis_ctx = context.Background()
)
// first service: get_users
func (s *userServiceServer) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	errRedise := client.Get(redis_ctx, req.GetAuthKey()).Err()
	if errRedise != nil {
		return nil, errors.New("invalid auth_key")
	}


	if req.GetMessageId()%2 != 0 || req.GetMessageId() <= 0 {
		return nil, errors.New("invalid message_id")
	}
	r, _ := regexp.Compile("\\D")
	if r.MatchString(req.GetUserId()){
		return nil, errors.New("invalid input type")
	}
	var rows *sql.Rows
	var err error

	// Check if user_id is empty
	if req.GetUserId() == "" {
		// Retrieve the first 10 rows from the table
		rows, err = s.db.Query("SELECT * FROM users LIMIT 10")
	} else {
		// Retrieve the row with the specified user_id
		rows, err = s.db.Query("SELECT * FROM users WHERE id = $1", req.GetUserId())
	}
	
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*pb.User

	for rows.Next() {
		var name, family string
    	var id, age int32
    	var sex string
    	var createdAt time.Time
		err := rows.Scan(&name, &family, &id, &age, &sex, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		user := &pb.User{
			Name:      name,
			Family:    family,
			Id:        id,
			Age:       age,
			Sex:       sex,
			CreatedAt: createdAt.String(),
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return &pb.GetUsersResponse{
		Users:     users,
		MessageId: req.MessageId + 1,
	}, nil
}

// second service: get_users_with_sql_inject
func (s *userServiceServer) GetUsersWithSQLInject(ctx context.Context, req *pb.GetUsersWithSQLInjectRequest) (*pb.GetUsersWithSQLInjectResponse, error) {
	errRedise := client.Get(redis_ctx, req.GetAuthKey()).Err()
	if errRedise != nil {
		return nil, errors.New("invalid auth_key")
	}

	if req.GetMessageId()%2 != 0 || req.GetMessageId() <= 0 {
		return nil, errors.New("invalid message_id")
	}
	
	var rows *sql.Rows
	var err error

	// Check if user_id is empty
	if req.GetUserId() == "" {
		// Retrieve the first 10 rows from the table
		rows, err = s.db.Query("SELECT * FROM users LIMIT 10")
	} else {
		// Retrieve the row with the specified user_id
		query := "SELECT * FROM users WHERE id = " + req.GetUserId()
		rows, err = s.db.Query(query)
	}
	
	
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []*pb.User

	for rows.Next() {
		var name, family string
    	var id, age int32
    	var sex string
    	var createdAt time.Time
		err := rows.Scan(&name, &family, &id, &age, &sex, &createdAt)
		if err != nil {
			log.Fatal(err)
		}
		user := &pb.User{
			Name:      name,
			Family:    family,
			Id:        id,
			Age:       age,
			Sex:       sex,
			CreatedAt: createdAt.String(),
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	return &pb.GetUsersWithSQLInjectResponse{
		Users:     users,
		MessageId: req.MessageId + 1,
	}, nil
}

func main() {
	db, err := sql.Open("postgres", "host=10.10.0.2 port=5432 user=postgres password=password dbname=database_0 sslmode=disable")
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := grpc.NewServer()
	userService := &userServiceServer{db: db}
	pb.RegisterUserServiceServer(server, userService) 

	listener, err := net.Listen("tcp", ":5062")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC server
	log.Println("Server started, listening on port 5062")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}