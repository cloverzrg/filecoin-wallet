package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var DB *gorm.DB

func Connect() (err error) {
	_, err = os.Stat("./data")
	if err == os.ErrNotExist {
		_, err := os.Create("./data")
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	DB, err = gorm.Open("sqlite3", "./data/sqlite3.db")
	if err != nil {
		return err
	}
	return err
}
