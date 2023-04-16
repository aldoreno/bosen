package domain

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Username  string
	UserModel struct {
		ID         uuid.UUID `gorm:"column:uid;primarykey"`
		Username   Username  `gorm:"column:username;"`
		Password   Password  `gorm:"column:password;"`
		FirstName  string    `gorm:"column:firstName;"`
		MiddleName string    `gorm:"column:middleName;"`
		LastName   string    `gorm:"column:lastName;"`
	}
)

func (u Username) String() string {
	return string(u)
}

func (u UserModel) TableName() string {
	return "userinfo"
}
