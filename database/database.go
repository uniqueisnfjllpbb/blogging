package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/uniqueisnfjllpbb/blogging/model"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"

	//"github.com/oklog/ulid"
	//"github.com/uniqueisnfjllpbb/blogging/model"
	"log"
)

var d *gorm.DB

func InitDB() *gorm.DB {
	errEnv := godotenv.Load(".env")
	db_host := os.Getenv("DB_HOST")
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	if errEnv != nil {
		fmt.Printf("読み込み出来ませんでした: %v", errEnv)
	}
	dbenv := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", db_host, db_username, db_password, db_database)
	db, err := gorm.Open(postgres.Open(dbenv), &gorm.Config{})
	//defer db.Close()

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("DB接続に成功しました。")
	}
	create := db.AutoMigrate(&model.Accounts{})

	if create != nil {
		log.Fatalln(create)
	} else {
		fmt.Println("DB作成に成功しました。")
	}
	//
	//t := time.Now()
	//entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	//id := ulid.MustNew(ulid.Timestamp(t), entropy)

	user := model.Accounts{

		FirstName: "testFirstname1",
		LastName:  "testLastname1",
		Email:     "test1@gmail.com",
		Password:  "passwordforuser1",
		IsActive:  true,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	result := db.Create(&user)

	if result != nil {
		log.Fatalln(result)
	} else {
		fmt.Println("DBにデータを挿入できました。")
	}

	return db
}
