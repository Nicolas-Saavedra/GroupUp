package model

type RatingTag struct {
	Base
	Value    string
	RatingID string
}

func (m *RatingTag) TableName() string {
	return "rating_tag"
}
