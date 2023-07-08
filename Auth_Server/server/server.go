package main

import (
	pb "HW1"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var (
	port                    = flag.Int("port", 5052, "The server port")
	p                uint32 = 23
	g                int32  = 5
	//redis
	client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
	redis_ctx = context.Background()
	user_snonce = make(map[string]string)
	user_msgid = make(map[string]uint32)
)

func random_str() (string){
	rand.Seed(time.Now().Unix() * 2)
    length := 20
  
    ran_str := make([]byte, length)
  
    // Generating Random string
    for i := 0; i < length; i++ {
        ran_str[i] = byte((65 + rand.Intn(25)))
    }
  
    return string(ran_str)
}

type server struct {
	pb.UnimplementedReq_DHServer
}




func (s *server) GetPQResponse(ctx context.Context, in *pb.PQ_Request) (*pb.PQ_Response, error) {
	if len(in.GetNonce()) != 20 {
		return nil, fmt.Errorf("given nonce doesn't have the required length")
	}

	if in.GetMessageId() <= 0 || in.GetMessageId()%2 != 0 {
		return nil, errors.New("given message_id doesn't have the correct form")
	}

	log.Printf("Received: %v", in.GetNonce())

	server_nonce := random_str()
	user_snonce[in.GetNonce()] = server_nonce

	user_msgid[in.GetNonce()] = in.GetMessageId()

	//SHA1 hash
	h := sha1.New()
	h.Write([]byte(in.GetNonce() + server_nonce))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	//redis
	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
	redis_ctx := context.Background()
	err := client.Set(redis_ctx, fmt.Sprintf("%v", sha1_hash), user_msgid[in.GetNonce()] + 1, 20*time.Minute).Err()
	if (err != nil){
		return nil, err
	}

	return &pb.PQ_Response{Nonce: in.GetNonce(), ServerNonce: server_nonce, MessageId: in.GetMessageId() + 1, P: p, G: g}, nil
}

func (s *server) GetDHResponse(ctx context.Context, in *pb.DH_Request) (*pb.DH_Response, error) {
	if len(in.GetNonce()) != 20 {
		return nil, fmt.Errorf("given nonce doesn't have the required length")
	}

	// if in.GetNonce() != nonce {
	// 	return nil, fmt.Errorf("given nonce is not valid")
	// }

	if len(in.GetServerNonce()) != 20 {
		return nil, fmt.Errorf("given server_nonce doesn't have the required length")
	}

	if in.GetServerNonce() != user_snonce[in.GetNonce()] {
		return nil, fmt.Errorf("given server_nonce is not valid")
	}

	if in.GetMessageId() <= 0 || in.GetMessageId()%2 != 0 {
		return nil, errors.New("given message_id doesn't have the correct form")
	}

	if in.GetMessageId() <= user_msgid[in.GetNonce()] {
		return nil, errors.New("given message_id is not valid")
	}

	log.Printf("Received: %v", in.GetNonce())
	b := 15
	B := math.Mod(math.Pow(float64(g), float64(b)), float64(p))
	public_key := math.Mod(math.Pow(float64(in.GetA()), float64(b)), float64(p))

	//SHA1 hash
	h := sha1.New()
	h.Write([]byte(in.GetNonce() + in.GetServerNonce()))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	
	client.Del(redis_ctx, sha1_hash).Err()
	err := client.Set(redis_ctx, fmt.Sprintf("%v", public_key), sha1_hash, 20*time.Minute).Err()
	if (err != nil){
		return nil, err
	}

	return &pb.DH_Response{Nonce: in.GetNonce(), ServerNonce: in.GetServerNonce(), MessageId: in.GetMessageId() + 1, B: int32(B)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterReq_DHServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve server: %v", err)
	}
}
