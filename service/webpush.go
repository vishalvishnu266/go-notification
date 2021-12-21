package service

import (
	"encoding/json"
	webpush "github.com/SherClockHolmes/webpush-go"
	"os"
)

func SendWebpush(subscription string, message string) {
	vapidPublicKey := os.Getenv("VAPIDPUBLIC")
	vapidPrivateKey := os.Getenv("VAPIDPRIVATE")
	s := webpush.Subscription{}
	json.Unmarshal([]byte(subscription), &s)
	resp, _ := webpush.SendNotification([]byte(message), &s, &webpush.Options{
		Subscriber:      "admin@growmax.com", // Do not include "mailto:"
		VAPIDPublicKey:  vapidPublicKey,
		VAPIDPrivateKey: vapidPrivateKey,
		TTL:             30,
	})
	defer resp.Body.Close()
	// return resp.Status, err
}
