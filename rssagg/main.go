package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Fatal("PORT is not found in the Environment")
	}

	fmt.Println("Port: ", portStr)
}
