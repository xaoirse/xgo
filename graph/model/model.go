package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatalln("Error in opening db:", err)
	}
	db.AutoMigrate(
		// &Session{},
		&User{})
	return db
}

// BeforeCreate is a gorm hook
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
