package Auth_Server

import (
	pb "HW1"
	"context"
	"flag"
	"log"
	"math"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:5052", "The address to connect to")
)

func random_str() string {
	rand.Seed(time.Now().Unix())
	length := 20

	ran_str := make([]byte, length)

	// Generating Random string
	for i := 0; i < length; i++ {
		ran_str[i] = byte((65 + rand.Intn(25)))
	}

	return string(ran_str)
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReq_DHClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var nonce string = random_str()
	var message_id uint32 = 2
	response, err := c.GetPQResponse(ctx, &pb.PQ_Request{Nonce: nonce, MessageId: message_id})
	if err != nil {
		log.Fatalf("Could not get PQ Response: %v", err)
	}

	a := 6
	A := math.Mod(math.Pow(float64(response.GetG()), float64(a)), float64(response.GetP()))
	response2, err := c.GetDHResponse(ctx, &pb.DH_Request{Nonce: nonce, ServerNonce: response.GetServerNonce(), MessageId: message_id + 2, A: int32(A)})
	if err != nil {
		log.Fatalf("Could not get DH Response: %v", err)
	}

	publicKey := math.Mod(float64(response.GetP()), math.Pow(float64(response2.GetB()), float64(a)))

	log.Print(publicKey)
}
