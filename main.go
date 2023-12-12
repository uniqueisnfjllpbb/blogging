package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/database"
)

func main() {
    r:= gin.Default()
    database.InitDB()

    
    r.Run(":8081")
    fmt.Println("localhost:8081")
    defer database.Close()

    r.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
}
