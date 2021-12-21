package service

import (
	"context"
	"log"

	"github.com/growmax/noti/apiproto"
)

func Centripush(channel string, notification string, client apiproto.CentrifugoApiClient) {
	message := `{"value": "` + notification + `"}`
	resp, err := client.Publish(context.Background(), &apiproto.PublishRequest{
		Channel: channel,
		Data:    []byte(message),
	})
	if err != nil {
		log.Printf("Transport level error: %v", err)
	} else {
		if resp.GetError() != nil {
			respError := resp.GetError()
			log.Printf("Error %d (%s)", respError.Code, respError.Message)
		}
	}
}
