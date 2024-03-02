package model

import (
	"gorm.io/gorm"
)

type RatingTag struct {
	gorm.Model
	Value    string
	RatingID uint
}

func (m *RatingTag) TableName() string {
	return "rating_tag"
}
