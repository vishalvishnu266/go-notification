package service

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReceiveFromKafka() {
	server := os.Getenv("KAFKASERVER")
	log.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"notification"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			// fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			// job := string(msg.Value)
			// saveJobToMongo(job)
		} else {
			// fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()

}
