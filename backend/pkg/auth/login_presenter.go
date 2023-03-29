package auth

import "bosen/pkg/domain"

type LoginPresenterImpl struct{}

func NewLoginPresenter() LoginPresenter {
	return LoginPresenterImpl{}
}

func (p LoginPresenterImpl) Output(token domain.Token) LoginOutput {
	return LoginOutput{
		Token: token.String(),
	}
}
