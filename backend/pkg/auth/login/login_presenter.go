package login

import (
	"bosen/application"
	"bosen/manifest"
	"bosen/pkg/domain"
	errs "bosen/pkg/errors"
	"context"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var _ LoginPresenter = (*LoginPresenterImpl)(nil)

type (
	LoginPresenterImpl struct {
		key string
	}
	TokenPayload struct {
		UserID     string `json:"user_id"`
		Username   string `json:"username"`
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
	}
)

func NewLoginPresenter(config application.Config) *LoginPresenterImpl {
	return &LoginPresenterImpl{key: config.JWTSecret}
}

func (p LoginPresenterImpl) Output(ctx context.Context, user *domain.UserModel) (*LoginOutput, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": manifest.AppName,
			"data": &TokenPayload{
				UserID:     user.ID.String(),
				Username:   user.Username.String(),
				FirstName:  user.FirstName,
				MiddleName: user.MiddleName,
				LastName:   user.LastName,
			},
		},
	)

	token, err := t.SignedString([]byte(p.key))
	if err != nil {
		zap.S().Errorf("unable to signed jwt: %w", err)
		return nil, errs.ErrAuthJwt
	}

	return &LoginOutput{Token: token}, nil
}
