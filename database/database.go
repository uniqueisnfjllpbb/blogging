package database

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"github.com/uniqueisnfjllpbb/blogging/model"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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
	create := db.AutoMigrate(&model.Accounts{}, &model.Post{}, &model.Profile{})
	if create != nil {
		log.Fatalln(create)
	} else {
		fmt.Println("DB作成に成功しました。")
	}
	return db
}

func InsertData(db *gorm.DB) {

	var accounts []model.Accounts
	for i := 0; i < 10; i++ {
		account := model.Accounts{
			//ID:        gofakeit.UUID(),
			Firstname: gofakeit.FirstName(),
			Lastname:  gofakeit.LastName(),
			Email:     gofakeit.Email(),
			Password:  gofakeit.Password(true, true, true, true, true, 32),
		}
		accounts = append(accounts, account)

	}

	fmt.Println(accounts)

	for _, account := range accounts {
		//fmt.Printf("ID: %s\n", account.ID)
		fmt.Printf("Firstname: %s\n", account.Firstname)
		fmt.Printf("Lastname: %s\n", account.Lastname)
		fmt.Printf("Email: %s\n\n", account.Email)
		fmt.Printf("Password: %s\n\n", account.Password)
	}
	var posts []model.Post
	for i := 0; i < 10; i++ {
		post := model.Post{
			//ID:        gofakeit.UUID(),
			Title: gofakeit.BookTitle(),
			Body:  gofakeit.LoremIpsumSentence(100),
		}
		posts = append(posts, post)
	}

	fmt.Println(posts)

	for _, post := range posts {
		//fmt.Printf("ID: %s\n", account.ID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Body: %s\n", post.Body)
	}

	resultaccount := db.Create(&accounts)
	resultpost := db.Create(&posts)

	if resultaccount != nil {
		log.Fatalln(resultaccount.Error)
	} else {
		fmt.Println("DBにアカウントデータを挿入できました。")
	}

	if resultpost != nil {
		log.Fatalln(resultpost.Error)
	} else {
		fmt.Println("DBにポストデータを挿入できました。")
	}

}
