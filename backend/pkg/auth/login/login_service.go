package login

import (
	"bosen/internal/trace"
	"bosen/pkg/domain"
	errs "bosen/pkg/errors"
	"bosen/pkg/runtime"
	userpkg "bosen/pkg/user"
	"context"
	"errors"
)

var _ LoginService = (*LoginServiceImpl)(nil)

type (
	LoginInput struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginPresenter interface {
		Output(context.Context, *domain.UserModel) (*LoginOutput, error)
	}

	LoginOutput struct {
		Token string `json:"token"`
	}

	LoginService interface {
		Login(context.Context, LoginInput) (*LoginOutput, error)
	}

	LoginServiceImpl struct {
		userRepo  userpkg.UserRepository
		presenter LoginPresenter
	}
)

func NewLoginServiceImpl(userRepo userpkg.UserRepository, presenter LoginPresenter) *LoginServiceImpl {
	return &LoginServiceImpl{userRepo, presenter}
}

func (s *LoginServiceImpl) Login(ctx context.Context, credentials LoginInput) (*LoginOutput, error) {
	span, ctx := trace.New(ctx, runtime.GetCurrentFunctionName())
	defer span.End()

	var user domain.UserModel

	criteria := userpkg.FindCriteria{
		Username: credentials.Username,
	}

	if err := s.userRepo.FindOne(ctx, criteria, &user); err != nil {
		switch {
		case errors.Is(err, errs.ErrUserNotFound):
			err = errs.ErrAuthCredentials
		}

		return nil, err
	}

	if !user.Password.CheckPasswordHash(credentials.Password) {
		return nil, errs.ErrAuthCredentials
	}

	return s.presenter.Output(ctx, &user)
}
