package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	DBConn *gorm.DB
)

func InitDataBase() {

	var err error
	DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	log.Print("Connection Opened to DataBase")
}
