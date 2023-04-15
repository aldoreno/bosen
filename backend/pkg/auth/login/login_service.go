package login

import (
	"bosen/pkg/domain"
	"bosen/pkg/errors"
	"bosen/pkg/user"
	"context"
	stderrors "errors"
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

func NewLoginServiceImpl(userRepo user.UserRepository, presenter LoginPresenter) *LoginServiceImpl {
	return &LoginServiceImpl{userRepo, presenter}
}

func (s *LoginServiceImpl) Login(ctx context.Context, credentials LoginInput) (LoginOutput, error) {
	var account user.User

	criteria := user.FindCriteria{
		Username: credentials.Username,
	}

	if err := s.userRepo.FindOne(ctx, criteria, &account); err != nil {
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			err = errors.ErrAuthCredentials
		}

		return s.presenter.Output(domain.Token{}), err
	}

	if account.Password != credentials.Password {
		return s.presenter.Output(domain.Token{}), errors.ErrAuthCredentials
	}

	token := domain.Token{
		Value: "123",
	}

	return s.presenter.Output(token), nil
}
