package main

import (
	"fdfd/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/registerUser", handlers.RegisterUser)

	log.Fatal(engine.Run("localhost:8082"))
}
