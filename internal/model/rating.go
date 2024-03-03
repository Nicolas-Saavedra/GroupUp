package model

type Rating struct {
	Base
	TextContent string
	StarRating  uint8
	UserID      string
	Tags        []RatingTag
}

func (m *Rating) TableName() string {
	return "rating"
}
