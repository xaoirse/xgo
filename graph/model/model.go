package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// GetModels return models
func GetModels() []interface{} {
	return []interface{}{
		&Session{},
		&User{},
	}
}

// BeforeCreate is a gorm hook
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
