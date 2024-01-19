package usecase

import (
	"fmt"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/authen/repositories"
)

type AuthUsecase struct {
	AccountRepositoty repositories.AccountRepositoty
}

func NewAuthUsecase(authRepository repositories.AccountRepositoty) *AuthUsecase {
	return &AuthUsecase{
		AccountRepositoty: authRepository,
	}
}

func (u *AuthUsecase) AuthenticateUser(account entities.LoginRequest) (entities.AccountRequest, bool, error) {
	user, err := u.AccountRepositoty.FindByUsername(account)
	if err != nil {
		return user, true, err
	}
	validate := user.Username != account.Username
	fmt.Println("validate :", validate)
	return user, validate, nil
}

func (u *AuthUsecase) Token(id int, access_token, refresh_token string) error {
	err := u.AccountRepositoty.Token(id, access_token, refresh_token)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecase) RefreshToken(refresh_token entities.RefreshTokenRequest) (entities.Token, error) {
	token, err := u.AccountRepositoty.FindRefreshToken(refresh_token)
	if err != nil {
		return token, err
	}
	return token, nil
}

func (u *AuthUsecase) UpdateRefreshToken(access_token, refresh_token string, id entities.Token) error {
	err := u.AccountRepositoty.UpdateToken(access_token, refresh_token, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthUsecase) DeleteToken(access_token, refresh_token string) error {
	err := u.AccountRepositoty.DeleteToken(access_token, refresh_token)
	if err != nil {
		return err
	}
	return nil
}
