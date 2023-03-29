package auth

import (
	"bosen/pkg/domain"
	"bosen/pkg/user"
	"context"
)

var _ LoginService = (*LoginServiceImpl)(nil)

type (
	LoginInput struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginPresenter interface {
		Output(domain.Token) LoginOutput
	}

	LoginOutput struct {
		Token string `json:"token"`
	}

	LoginService interface {
		Login(context.Context, LoginInput) (LoginOutput, error)
	}

	LoginServiceImpl struct {
		userRepo  user.UserRepository
		presenter LoginPresenter
	}
)

func NewAuthServiceImpl(userRepo user.UserRepository) *LoginServiceImpl {
	return &LoginServiceImpl{
		userRepo: userRepo,
	}
}

func (s *LoginServiceImpl) Login(ctx context.Context, credentials LoginInput) (LoginOutput, error) {
	var account user.User

	criteria := user.FindCriteria{
		Username: credentials.Username,
	}

	if err := s.userRepo.FindOne(ctx, criteria, &account); err != nil {
		return s.presenter.Output(domain.Token{}), ErrAccountNotFound
	}

	if credentials.Password != "password" {
		return s.presenter.Output(domain.Token{}), ErrWrongUsernameOrPassword
	}

	token := domain.Token{}

	return s.presenter.Output(token), nil
}