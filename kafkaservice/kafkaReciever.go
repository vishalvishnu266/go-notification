package kafkaservice

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func ReceiveFromKafka() {
	topic := "notification"
	brokerAddress := os.Getenv("KAFKASERVER")
	ctx := context.Background()

	// l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "notification",
		// Logger: l,
	})
	for {
		// msg, err := r.ReadMessage(ctx)
		// if err != nil {
		// 	log.Printf("could not read message " + err.Error())
		// }
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		// if err := r.CommitMessages(ctx, m); err != nil {
		// 	log.Fatal("failed to commit messages:", err)
		// }
		go KafkaNotiProcessor(r, m)
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
