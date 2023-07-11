package main

import (
	pb "HW1/proto"
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	_ "strconv"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "HW1/gateway/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	addr       = flag.String("addr", "localhost:5052", "The address to connect to")
	gateway_ctx context.Context
	c          pb.Req_DHClient
	P uint32
	G int32
	biz_client pb.UserServiceClient
)

type Blocklist struct {
	ips map[string]struct{}
	mu  sync.Mutex
}

type req_pq_send_data struct {
	ID uint32
}
type req_dh_params struct {
	NONCE        string
	SERVER_NONCE string
	ID           uint32
	KEY          int32
}
type get_users_biz struct {
	USER_ID    string
	AUTH_KEY   string
	MESSAGE_ID int32
}
type get_users_injection_biz struct {
	USER_ID    string
	AUTH_KEY   string
	MESSAGE_ID int32
}

func NewBlocklist() *Blocklist {
	return &Blocklist{
		ips: make(map[string]struct{}),
	}
}

func (bl *Blocklist) BlockIP(ip string) {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	bl.ips[ip] = struct{}{}
	time.AfterFunc(24*time.Hour, func() {
		bl.mu.Lock()
		defer bl.mu.Unlock()
		delete(bl.ips, ip)
	})
}

func (bl *Blocklist) IsBlocked(ip string) bool {
	bl.mu.Lock()
	defer bl.mu.Unlock()
	_, blocked := bl.ips[ip]
	return blocked
}
func RateLimiter(limit rate.Limit, burst int, blocklist *Blocklist) gin.HandlerFunc {
	limiter := rate.NewLimiter(limit, burst)

	return func(c *gin.Context) {
		ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)

		if blocklist.IsBlocked(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "You have been blocked due to excessive requests.",
			})
			c.Abort()
			return
		}

		if !limiter.Allow() {
			blocklist.BlockIP(ip)
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
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

// @Summary Request PQ
// @Description Request PQ from the server
// @Accept json
// @Produce json
// @Param data body req_pq_send_data true "Request data"
// @Success 200 {object} gin.H
// @Failure 400 {string} string "Bad Request"
// @Router /auth/req_pq/ [post]
func req_pq(ctx *gin.Context){
	var nonce string = random_str()
	var data req_pq_send_data
	if ctx.BindJSON(&data) != nil {
		ctx.JSON(200, gin.H{
			"err": "wrong json format!",
		})
	} else {
		response, err := c.GetPQResponse(gateway_ctx, &pb.PQ_Request{Nonce: nonce, MessageId: data.ID})
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
		} else {
			P = response.GetP()
			G = response.GetG()
			ctx.JSON(200, gin.H{
				"nonce":       response.GetNonce(),
				"serverNonce": response.GetServerNonce(),
				"message_id":  response.GetMessageId(),
				"p":           response.GetP(),
				"g":           response.GetG(),
			})
		}
	}
}

// @Summary Request DH Params
// @Description Request DH Params from the server
// @Accept json
// @Produce json
// @Param dh_params body req_dh_params true "Request data"
// @Success 200 {object} gin.H
// @Failure 400 {string} string "Bad Request"
// @Router /auth/req_DH_params/ [post]
func req_dh(ctx *gin.Context) {
	var dh_params req_dh_params

	if ctx.BindJSON(&dh_params) != nil {
		ctx.JSON(200, gin.H{
			"err": "wrong json format!",
		})
	} else {
		A := math.Mod(math.Pow(float64(G), float64(dh_params.KEY)), float64(P))
		fmt.Println(A)
		response, err := c.GetDHResponse(gateway_ctx, &pb.DH_Request{Nonce: dh_params.NONCE,
			ServerNonce: dh_params.SERVER_NONCE, MessageId: dh_params.ID, A: int32(A)})
		if err != nil {
			fmt.Println(err)
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
		} else {

			key := math.Mod(math.Pow(float64(response.GetB()), float64(dh_params.KEY)), float64(P))
			ctx.JSON(200, gin.H{
				"B":   response.GetB(),
				"key": key,
			})
		}
	}
}

// @Summary Get Users
// @Description Get users from the server
// @Accept json
// @Produce json
// @Param data body get_users_biz true "Request data"
// @Success 200 {object} gin.H
// @Failure 400 {string} string "Bad Request"
// @Router /biz/getUsers/ [post]
func get_users(ctx *gin.Context) {
	var data get_users_biz
	ctx.BindJSON(&data)
	fmt.Print(data.USER_ID)
	response, err := biz_client.GetUsers(gateway_ctx, &pb.GetUsersRequest{UserId: data.USER_ID, AuthKey: data.AUTH_KEY, MessageId: data.MESSAGE_ID})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, gin.H{
			"err": err.Error(),
		})
	} else {
		var s string
		for _, user := range response.Users {
			s += fmt.Sprint(user.Id) + "  " + user.Name + "  " + user.Family + "  " + fmt.Sprint(user.Age) + "  " + user.Sex + "  " + user.CreatedAt + "\u000a"
		}
		ctx.JSON(200, gin.H{
			"data":       s,
			"message_id": response.MessageId,
		})
	}
}


// @Summary Get Users with Injection
// @Description Get users from the server with SQL injection
// @Accept json
// @Produce json
// @Param data body get_users_injection_biz true "Request data"
// @Success 200 {object} gin.H
// @Failure 400 {string} string "Bad Request"
// @Router /biz/WithInjection/ [post]
func get_users_with_injection(ctx *gin.Context) {
	var data get_users_injection_biz
	if ctx.BindJSON(&data) != nil {
		fmt.Printf(data.USER_ID)
		ctx.JSON(200, gin.H{
			"err": "wrong json format!",
		})
	} else {
		response, err := biz_client.GetUsersWithSQLInject(gateway_ctx, &pb.GetUsersWithSQLInjectRequest{
			UserId: data.USER_ID, AuthKey: data.AUTH_KEY, MessageId: data.MESSAGE_ID,
		})
		if err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
		} else {
			var s string
			for _, user := range response.Users {
				s += fmt.Sprint(user.Id) + "  " + user.Name + "  " + user.Family + "  " + fmt.Sprint(user.Age) + "  " + user.Sex + "  " + user.CreatedAt + "\u000a"
			}
			ctx.JSON(200, gin.H{
				"data is":       s,
				"message_id is": response.MessageId,
			})
		}
	}
}

// @title Gateway API
// @version 1.0
// @description A gateway to our web-service
func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewReq_DHClient(conn)
	///////////////////////////////////////
	conn_biz, err_conn_biz := grpc.Dial("localhost:5062", grpc.WithInsecure())
	if err_conn_biz != nil {
		log.Fatalf("Failed to connect to bizServer! : %v", err_conn_biz)
	}
	defer conn_biz.Close()
	biz_client = pb.NewUserServiceClient(conn_biz)

	gateway_ctx = context.Background()
	
	// Create the blocklist and rate limiter
	blocklist := NewBlocklist()
	rateLimiter := RateLimiter(rate.Limit(100), 10, blocklist)

	// Create the Gin router
	router := gin.Default()
	router.Use(cors.Default())

	/////////////////////////// POSTS /////////////////////////////////////////////

	router.POST("/auth/req_pq/", rateLimiter, req_pq)

	router.POST("/auth/req_DH_params/", rateLimiter, req_dh)

	router.POST("/biz/getUsers/", get_users)

	router.POST("/biz/WithInjection/", get_users_with_injection)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err2 := router.Run(":6433")
	if err2 != nil {
		log.Fatal("Failed to start the server:", err2)
	}
}
