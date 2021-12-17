package main

import (
	"log"
	"os"

	"github.com/growmax/noti/db"
	"github.com/growmax/noti/router"
	// "github.com/growmax/noti/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = db.GetMongoClient()
	if err != nil {
		log.Printf("Error while connecting to MongoDB")
	} else {
		log.Printf("connected to mongodb successfully!!!!!!!!!!!!")
	}
	// service.ReceiveFromKafka()
	r := router.SetupRouter()
	port := os.Getenv("PORT")
	r.Run(port)
}
