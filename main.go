package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uniqueisnfjllpbb/blogging/database"
	"github.com/uniqueisnfjllpbb/blogging/router"
)

func main() {
	r := gin.Default()
	router.Routes(r)
	dbCon := database.InitDB()
	database.InsertData(dbCon)
	r.Run(":8081")
}
