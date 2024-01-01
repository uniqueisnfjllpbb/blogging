package database

import (
	"database/sql"
	"fmt"
	//"github.com/oklog/ulid"
	//"github.com/uniqueisnfjllpbb/blogging/model"
	"log"
	//"math/rand"
	"os"
	//"time"

	// "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ←これを追記
	_ "gorm.io/driver/postgres"
)

var d *sql.DB

func InitDB() *sql.DB {
	errEnv := godotenv.Load(".env")
	db_host := os.Getenv("DB_HOST")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	if errEnv != nil {
		fmt.Printf("読み込み出来ませんでした: %v", errEnv)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("%s:%s@%s:5432/%s?sslmode=disable", db_username, db_password, db_host, db_database))
	defer db.Close()

	// postgres://
	// "admin:passwordpassword@/blg_general?sslmode=disable"
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("DB接続に成功しました。")
	}

	//user := model.User{
	//	ID:        id,
	//	FirstName: "testFirstname1",
	//	LastName:  "test1@gmail.com",
	//	Password:  "passwordforuser1",
	//	IsActive:  true,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//}

	return db
}
