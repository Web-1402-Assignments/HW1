package main

import (
	pb "Auth_Server"
	"context"
	"crypto/sha1"
	"encoding/hex"
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
	nonce            string
	server_nonce     string
	first_message_id uint32
)

func random_str() (string){
	rand.Seed(time.Now().Unix())
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
		return nil, fmt.Errorf("given message_id doesn't have the correct form")
	}

	log.Printf("Received: %v", in.GetNonce())
	nonce = in.GetNonce()
	server_nonce = random_str()
	first_message_id = in.GetMessageId()

	return &pb.PQ_Response{Nonce: nonce, ServerNonce: server_nonce, MessageId: in.MessageId + 1, P: p, G: g}, nil
}

func (s *server) GetDHResponse(ctx context.Context, in *pb.DH_Request) (*pb.DH_Response, error) {
	if len(in.GetNonce()) != 20 {
		return nil, fmt.Errorf("given nonce doesn't have the required length")
	}

	if in.GetNonce() != nonce {
		return nil, fmt.Errorf("given nonce is not valid")
	}

	if len(in.GetServerNonce()) != 20 {
		return nil, fmt.Errorf("given server_nonce doesn't have the required length")
	}

	if in.GetServerNonce() != server_nonce {
		return nil, fmt.Errorf("given server_nonce is not valid")
	}

	if in.GetMessageId() <= 0 || in.GetMessageId()%2 != 0 {
		return nil, fmt.Errorf("given message_id doesn't have the correct form")
	}

	if in.GetMessageId() <= first_message_id {
		return nil, fmt.Errorf("given message_id is not valid")
	}

	log.Printf("Received: %v", in.GetNonce())
	b := 15
	B := math.Mod(math.Pow(float64(g), float64(b)), float64(p))
	public_key := math.Mod(float64(p), math.Pow(float64(in.GetA()), float64(b)))

	//SHA1 hash
	h := sha1.New()
	h.Write([]byte(nonce + server_nonce))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	//redis
	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
	redis_ctx := context.Background()
	err := client.Set(redis_ctx, fmt.Sprintf("%v", public_key), sha1_hash, 0)
	if (err != nil){
		return nil, fmt.Errorf(err.Result())
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
