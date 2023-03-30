package user

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		ID         uuid.UUID `gorm:"column:uid;primarykey"`
		Username   string    `gorm:"column:username;"`
		Password   string    `gorm:"column:password;"`
		FirstName  string    `gorm:"column:firstName;"`
		MiddleName string    `gorm:"column:middleName;"`
		LastName   string    `gorm:"column:lastName;"`
	}
)

func (u User) TableName() string {
	return "userinfo"
}
