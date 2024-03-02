package model

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Name   string
	Groups []Group
}

func (m *Course) TableName() string {
	return "course"
}
