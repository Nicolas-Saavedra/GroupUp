package model

type Rating struct {
	Base
	TextContent string
	StarRating  uint8
	UserID      uint
	Tags        []RatingTag
}

func (m *Rating) TableName() string {
	return "rating"
}
