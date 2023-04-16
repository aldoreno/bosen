package login

import (
	"bosen/pkg/domain"
	"bosen/pkg/errors"
	userpkg "bosen/pkg/user"
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
		userRepo  userpkg.UserRepository
		presenter LoginPresenter
	}
)

func NewLoginServiceImpl(userRepo userpkg.UserRepository, presenter LoginPresenter) *LoginServiceImpl {
	return &LoginServiceImpl{userRepo, presenter}
}

func (s *LoginServiceImpl) Login(ctx context.Context, credentials LoginInput) (LoginOutput, error) {
	var user domain.UserModel

	criteria := userpkg.FindCriteria{
		Username: credentials.Username,
	}

	if err := s.userRepo.FindOne(ctx, criteria, &user); err != nil {
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			err = errors.ErrAuthCredentials
		}

		return s.presenter.Output(domain.Token{}), err
	}

	if !user.Password.CheckPasswordHash(credentials.Password) {
		return s.presenter.Output(domain.Token{}), errors.ErrAuthCredentials
	}

	token := domain.Token{
		Value: "123",
	}

	return s.presenter.Output(token), nil
}
