package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ←これを追記
     "gorm.io/driver/postgres"
)

var d *gorm.DB

// type User struct {
//     gorm.Model
//     Name  string
//     Email string
// }


func InitDB() *gorm.DB {
    errEnv := godotenv.Load(".env")
    db_host := os.Getenv("DB_HOST")
    db_database := os.Getenv("DB_DATABASE")
    db_username := os.Getenv("DB_USERNAME")
    db_password := os.Getenv("DB_PASSWORD")
    if errEnv != nil {
		fmt.Printf("読み込み出来ませんでした: %v", errEnv)
	} 
    db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",db_username, db_password, db_host, db_database))
     
    // "admin:passwordpassword@/blg_general?sslmode=disable"
    if err != nil {
        log.Fatalln(err)
    }



    return db
}



// GetDB returns database connection
func GetDB() *gorm.DB {
    return d
}

// Close closes database
func Close() {
    d.Close()
}