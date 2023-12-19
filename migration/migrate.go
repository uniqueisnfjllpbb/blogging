package main

import (
	"fmt"
	"github.com/uniqueisnfjllpbb/blogging/db"
	"github.com/uniqueisnfjllpbb/blogging/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("マイグレーション成功")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
