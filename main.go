package main

import (
	"fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/database"
)

func main() {
	r := gin.Default()
	database.InitDB()

	fmt.Println("localhost:8081")
	defer database.Close()

	r.Run(":8081")
}
