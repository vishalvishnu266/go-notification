package main

import (
	"fmt"
	"log"
	"os"

	"github.com/growmax/noti/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	password := os.Getenv("PASSWORD")
	fmt.Print(password)
	r := router.SetupRouter()
	fmt.Print()
	r.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
