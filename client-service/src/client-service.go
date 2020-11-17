package main

import (
	"context"
    "fmt"
    "log"
	"os"
	"time"
    "./proto/messagepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
    "github.com/gin-gonic/gin"
)

// 環境変数からAPIサービスのエンドポイント取得
var (
    ApiServiceURL      = os.Getenv("API_SERVICE_URL")
    ApiServicePort     = os.Getenv("API_SERVICE_PORT")
    ApiRequestProtocol = os.Getenv("API_REQUEST_PROTOCOL")
)

type ResponseBody struct {
    Text string `json:"text"`
}

func WrapperFunc(fn func(c *gin.Context)) gin.HandlerFunc {
    return func(c *gin.Context) {
        fn(c)
    }
}

// APIサービスにリクエストしてレスポンス作成
func CreateResponse(c *gin.Context) {
	fmt.Println("Hello. I'm a client service")
	headerText := c.Query("header")
	fmt.Println("header: ", headerText)
	start := time.Now()

	url := ApiServiceURL + ":" + ApiServicePort
	fmt.Println(url)

	var opts grpc.DialOption
	var conn *grpc.ClientConn
	var err error
	var res *messagepb.MessageResponse

	if ApiRequestProtocol == "http" {
		opts = grpc.WithInsecure()
	} else {
		// HTTPSでgrpc通信を行う場合、オレオレ証明書が必要
		creds, err := credentials.NewClientTLSFromFile("/etc/ssl/certs/ca-certificates.crt", "")
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
		opts = grpc.WithTransportCredentials(creds)
	}
	// 接続
	conn, err = grpc.Dial(url, opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// defer conn.Close()
	client := messagepb.NewMessageServiceClient(conn)

	req := &messagepb.MessageRequest{Message: headerText}
	// 接続 (100回)
	for i := 0; i < 100; i++ {
		res, err = client.Message(context.Background(), req)
	}
	// res, err := client.Message(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not response: %v", err)
    }
    returnValue := res.Message + " world!"
    response := ResponseBody{
        Text: returnValue,
	}
	
	end := time.Now()
	fmt.Printf("%f\n", (end.Sub(start)).Seconds())
    c.JSON(200, response)
}

func main() {
    r := gin.Default()
    r.GET("/messages", WrapperFunc(CreateResponse))
    r.Run(":8090")
}