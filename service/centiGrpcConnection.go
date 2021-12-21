package service

import (
	"context"
	"log"
	"os"
	"sync"

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

var centriOnce sync.Once

var client *apiproto.CentrifugoApiClient

func GetCentriConnection() *apiproto.CentrifugoApiClient {
	centriOnce.Do(func() {

		addr := os.Getenv("GRPC_SERVER")
		key := os.Getenv("GRPC_KEY")
		conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithPerRPCCredentials(keyAuth{key}))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Connect to grpc **************************")
		// defer conn.Close()
		c := apiproto.NewCentrifugoApiClient(conn)
		client = &c
	})
	return client

}
