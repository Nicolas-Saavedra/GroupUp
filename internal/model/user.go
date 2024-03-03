package model

type User struct {
	BaseUser
	Name    string    `gorm:"not null"`
	Email   string    `gorm:"unique;not null"`
	Groups  []Group   `gorm:"many2many:user_groups" json:",omitempty"`
	Tags    []UserTag `                             json:",omitempty"`
	Ratings []Rating  `                             json:",omitempty"`
}

func (u *User) TableName() string {
	return "users"
}
