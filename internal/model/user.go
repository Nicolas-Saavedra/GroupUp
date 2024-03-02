package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"not null"`
	Password string  `gorm:"not null"              json:"-"`
	Email    string  `gorm:"unique;not null"`
	Groups   []Group `gorm:"many2many:user_groups"`
	Tags     []UserTag
	Ratings  []Rating
}

func (u *User) TableName() string {
	return "users"
}
