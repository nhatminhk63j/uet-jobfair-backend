package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("sqlite3", "uet_job_fair.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Company{})

	DB = db
}
