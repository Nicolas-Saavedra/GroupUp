package model

import (
	"gorm.io/gorm"
)

type UserTag struct {
	gorm.Model
	Value  string
	UserID uint
}

func (m *UserTag) TableName() string {
	return "user_tag"
}
