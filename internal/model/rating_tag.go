package model

type RatingTag struct {
	Base
	Value    string
	RatingID uint
}

func (m *RatingTag) TableName() string {
	return "rating_tag"
}
