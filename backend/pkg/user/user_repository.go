package user

import (
	"context"

	"gorm.io/gorm"
)

var _ UserRepository = (*UserRepositoryImpl)(nil)

type (
	UserRepository interface {
		FindOne(context.Context, FindCriteria, *User) error
	}

	UserRepositoryImpl struct {
		db *gorm.DB
	}

	FindCriteria struct {
		Username string
	}

	User struct {
		gorm.Model
		ID         uint
		Username   string
		Password   string
		FirstName  string
		MiddleName string
		LastName   string
	}
)

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) FindOne(_ context.Context, criteria FindCriteria, user *User) error {
	var user_ User
	r.db.Model(&user_).Find(&user_)
	return nil
}
