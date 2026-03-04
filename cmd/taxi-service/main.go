package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := router.Run(":8080"); err != nil {
		log.Fatal("fail run server on port 8080:", err)
	}
}
