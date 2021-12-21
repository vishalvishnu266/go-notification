package kafkaservice

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/growmax/noti/db"
	"github.com/growmax/noti/model"
	"github.com/growmax/noti/service"
	"github.com/segmentio/kafka-go"
)

func KafkaNotiProcessor(r *kafka.Reader, msg kafka.Message) {
	var s model.KafkaData
	err := json.Unmarshal([]byte(msg.Value), &s)
	if err != nil {
		log.Println(err)
	}
	var noti model.Notification
	noti.User = s.To
	noti.Notification = s.Message
	noti.CreatedAt = time.Now()
	err = db.AddNotification(noti)
	if err != nil {
		return
	}
	service.SendNotification(noti)

	if err := r.CommitMessages(context.Background(), msg); err != nil {
		log.Fatal("failed to commit messages:", err)
	}
}
