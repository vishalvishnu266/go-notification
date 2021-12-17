package service

import (
	"encoding/json"
	"os"

	webpush "github.com/SherClockHolmes/webpush-go"
)

func notification(subscription string, message string) (string, error) {

	vapidPublicKey := os.Getenv("VAPIDPUBLIC")
	vapidPrivateKey := os.Getenv("VAPIDPRIVATE")
	s := webpush.Subscription{}

	json.Unmarshal([]byte(subscription), &s)
	resp, err := webpush.SendNotification([]byte(message), &s, &webpush.Options{
		Subscriber:      "example@example.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	defer resp.Body.Close()
	return resp.Status, err
}
