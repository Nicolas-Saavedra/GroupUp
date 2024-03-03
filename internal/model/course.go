package model

type Course struct {
	Base
	Name   string
	Groups []Group
}

func (m *Course) TableName() string {
	return "course"
}
