package model

type UserTag struct {
	Base
	Value  string
	UserID string `json:"-"`
}

func (m *UserTag) TableName() string {
	return "user_tag"
}
