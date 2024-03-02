package model

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	TextContent string
	StarRating  uint8
	UserID      uint
	Tags        []RatingTag
}

func (m *Rating) TableName() string {
	return "rating"
}
