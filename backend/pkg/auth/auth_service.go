package auth

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Login(LoginInput) (*AuthToken, error)
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (a *authService) Login(input LoginInput) (*AuthToken, error) {
	if input.Username != "leo" {
		return nil, ErrAccountNotFound
	}

	if input.Password != "password" {
		return nil, ErrWrongUsernameOrPassword
	}

	token := &AuthToken{Token: "123"}
	return token, nil
}
