package model

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name     string `gorm:"not null"`
	MaxSize  uint   `gorm:"not null"`
	CourseID uint
	Members  []User `gorm:"many2many:user_groups"`
}

func (m *Group) TableName() string {
	return "group"
}
