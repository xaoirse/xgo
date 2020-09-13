// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewUser struct {
	Name string `json:"name"`
	// gorm.Model
	CreatedAt time.Time  ``
	UpdatedAt time.Time  ``
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	// gorm.Model
	CreatedAt time.Time  ``
	UpdatedAt time.Time  ``
	DeletedAt *time.Time `sql:"index"`
}
