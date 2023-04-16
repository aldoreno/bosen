package user

import uuid "github.com/satori/go.uuid"

type (
	UserModel struct {
		ID         uuid.UUID `gorm:"column:uid;primarykey"`
		Username   string    `gorm:"column:username;"`
		Password   string    `gorm:"column:password;"`
		FirstName  string    `gorm:"column:firstName;"`
		MiddleName string    `gorm:"column:middleName;"`
		LastName   string    `gorm:"column:lastName;"`
	}
)

func (u UserModel) TableName() string {
	return "userinfo"
}
