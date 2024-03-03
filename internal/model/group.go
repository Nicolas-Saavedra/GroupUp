package model

type Group struct {
	Base
	Name     string `gorm:"not null"`
	MaxSize  uint   `gorm:"not null"`
	CourseID uint
	Members  []User `gorm:"many2many:user_groups"`
}

func (m *Group) TableName() string {
	return "group"
}
