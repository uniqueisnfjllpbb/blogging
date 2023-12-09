package database

import (
	
	"log"

	"github.com/jinzhu/gorm"
    _ "github.com/lib/pq" // ←これを追記
)

var d *gorm.DB

// type User struct {
//     gorm.Model
//     Name  string
//     Email string
// }


func InitDB() *gorm.DB {
    

    db, err := gorm.Open("postgres", `postgres://admin:@localhost/blg_general?sslmode=disable`)
    
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