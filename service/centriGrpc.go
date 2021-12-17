package service

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/growmax/noti/apiproto"

	"google.golang.org/grpc"
)
type keyAuth struct {
    key string
}

func (t keyAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
    return map[string]string{
        "authorization": "apikey " + t.key,
    }, nil
}

func (t keyAuth) RequireTransportSecurity() bool {
    return false
}

func centripush() {
	key:= os.Getenv("GRPC_KEY")
    conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithPerRPCCredentials(keyAuth{key}))
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()
    client := apiproto.NewCentrifugoApiClient(conn)
    for {
        resp, err := client.Publish(context.Background(), &apiproto.PublishRequest{
            Channel: "chat:index",
            Data:    []byte(`{"input": "hello from GRPC"}`),
        })
        if err != nil {
            log.Printf("Transport level error: %v", err)
        } else {
            if resp.GetError() != nil {
                respError := resp.GetError()
                log.Printf("Error %d (%s)", respError.Code, respError.Message)
            } else {
                log.Println("Successfully published")
            }
        }
        time.Sleep(time.Second)
    }
}