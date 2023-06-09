package main

import (
	pb "Auth_Server"
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"

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

type server struct {
	pb.UnimplementedReq_DHServer
}

func (s *server) GetPQResponse(ctx context.Context, in *pb.PQ_Request) (*pb.PQ_Response, error) {
	if len(in.GetNonce()) != 20 {
		return nil, fmt.Errorf("Given Nonce doesn't have the required length")
	}

	if in.GetMessageId() <= 0 || in.GetMessageId()%2 != 0 {
		return nil, fmt.Errorf("Given Message_id doesn't have the correct form")
	}

	log.Printf("Received: %v", in.GetNonce())
	nonce = in.GetNonce()
	first_message_id = in.GetMessageId()

	return &pb.PQ_Response{Nonce: nonce, ServerNonce: server_nonce, MessageId: in.MessageId + 1, P: p, G: g}, nil
}

func (s *server) GetDHResponse(ctx context.Context, in *pb.DH_Request) (*pb.DH_Response, error) {
	if len(in.GetNonce()) != 20 {
		return nil, fmt.Errorf("Given Nonce doesn't have the required length")
	}

	if in.GetNonce() != nonce {
		return nil, fmt.Errorf("Given Nonce is not valid")
	}

	if len(in.GetServerNonce()) != 20 {
		return nil, fmt.Errorf("Given Server_Nonce doesn't have the required length")
	}

	if in.GetServerNonce() != server_nonce {
		return nil, fmt.Errorf("Given Server_Nonce is not valid")
	}

	if in.GetMessageId() <= 0 || in.GetMessageId()%2 != 0 {
		return nil, fmt.Errorf("Given Message_id doesn't have the correct form")
	}

	if in.GetMessageId() <= first_message_id {
		return nil, fmt.Errorf("Given Message_id is not valid")
	}

	log.Printf("Received: %v", in.GetNonce())
	b := 15
	B := math.Mod(math.Pow(float64(g), float64(b)), float64(p))
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
	log.Printf("PQserver listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve PQserver: %v", err)
	}
}
