package main

import (
	"fmt"
	"github.com/uniqueisnfjllpbb/blogging/fakedata"

	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/database"
)

func main() {
	r := gin.Default()
	database.InitDB()

	fmt.Println("localhost:8081")

	fakedata.FakeDataGenerator()
	defer database.Close()

	r.Run(":8081")
}
