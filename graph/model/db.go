package model

import (
	"log"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	// init sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetDb creates a database
func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.AutoMigrate(

		/**************************
		* Insert your models below *
		**************************/
		&User{},
	)
	return db
}

// BeforeCreate set uuid
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
