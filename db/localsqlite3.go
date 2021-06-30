package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB *gorm.DB

func Connect() (err error) {
	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		err := os.Mkdir("./data", os.ModePerm)
		if err != nil {
			return err
		}
	}
	DB, err = gorm.Open("sqlite3", "./data/sqlite3.db")
	DB = DB.Debug()
	if err != nil {
		return err
	}
	return err
}
