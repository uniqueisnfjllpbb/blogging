package database

import (
	"fmt"
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
	//
	//t := time.Now()
	//entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	//id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return db
}

func InsertData(db *gorm.DB) {

	user := model.Accounts{

		Firstname: "testFirstname1",
		Lastname:  "testLastname1",
		Email:     "test1@gmail.com",
		Password:  "passwordforuser1",
		Isactive:  true,
	}
	result := db.Create(&user)

	if result != nil {
		log.Fatalln(result.Error)
	} else {
		fmt.Println("DBにデータを挿入できました。")
	}

	fmt.Println("count:", result.RowsAffected)

}
