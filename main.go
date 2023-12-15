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

    
    
    fmt.Println("localhost:8081")
    defer database.Close()

    routeblog := r.Group("/blog")

    routeblog.GET("/all", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })

    r.Run(":8081")
}
