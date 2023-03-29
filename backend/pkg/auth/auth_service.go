package auth

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Login(LoginInput, *AuthToken) error
}

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (a *authService) Login(credentials LoginInput, token *AuthToken) error {
	if credentials.Username != "leo" {
		return ErrAccountNotFound
	}

	if credentials.Password != "password" {
		return ErrWrongUsernameOrPassword
	}

	*token = AuthToken{Token: "123"}

	return nil
}
