package model

type Group struct {
	Base
	Name     string `gorm:"not null"`
	MaxSize  uint   `gorm:"not null"`
	CourseID string
	Members  []User `gorm:"many2many:user_group"`
}

func (m *Group) TableName() string {
	return "group"
}
