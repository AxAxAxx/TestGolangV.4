package usecase

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/account/repositories"
	"github.com/AxAxAxx/go-test-api/pkg/middleware"
)

type AccountUsecase struct {
	UserRepositoty repositories.AccountRepositoty
}

func NewAccountUsecase(userRepositoty repositories.AccountRepositoty) *AccountUsecase {
	return &AccountUsecase{
		UserRepositoty: userRepositoty,
	}
}

func (u *AccountUsecase) CreateAccount(newAccount entities.RegisterAccount) error {
	hashpassword := newAccount.Password
	password, err := middleware.HashPassword(hashpassword)
	if err != nil {
		return err
	}
	account_id, err := u.UserRepositoty.CreateAccount(password, newAccount)
	if err != nil {
		return err
	}
	user_id, err := u.UserRepositoty.CreateUser(account_id, newAccount)
	if err != nil {
		return err
	}
	province_id, err := u.UserRepositoty.GetProvince(newAccount)
	if err != nil {
		return err
	}
	err = u.UserRepositoty.AddAddress(user_id, province_id, newAccount)
	if err != nil {
		return err
	}
	return nil
}
