package repositories

import (
	"time"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoty struct {
	DB *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepositoty {
	return &AccountRepositoty{
		DB: db,
	}
}

func (r *AccountRepositoty) CreateAccount(password string, newAccount entities.RegisterAccount) (int, error) {
	newAccount.CratedAt = time.Now()
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	query := `INSERT INTO public.account(
				username, password, crated_at)
				VALUES ($1, $2, $3) RETURNING account_id;`
	err = r.DB.QueryRowx(query, newAccount.Username, password, newAccount.CratedAt).Scan(&newAccount.AccountID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, nil
	}
	return newAccount.AccountID, nil
}

func (r *AccountRepositoty) CreateUser(account_id int, newUser entities.RegisterAccount) (int, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	insertuser := `INSERT INTO public."user"(
					first_name, last_name, email, dete_of_birth, phonenumber, account_id)
					VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id;`
	err = r.DB.QueryRowx(insertuser, newUser.FirstName, newUser.LastName, newUser.Email, newUser.DateOfBirth, newUser.PhoneNumber, account_id).Scan(&newUser.UserID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, nil
	}
	return newUser.UserID, nil
}

func (r *AccountRepositoty) GetProvince(getprovice entities.RegisterAccount) (int, error) {
	getprovince := `SELECT province_id FROM public.province WHERE province_name = $1;`
	err := r.DB.QueryRowx(getprovince, getprovice.Province).Scan(&getprovice.ProvinceID)
	if err != nil {
		return 0, err
	}
	return getprovice.ProvinceID, nil
}

func (r *AccountRepositoty) AddAddress(user_id, province_id int, addAddress entities.RegisterAccount) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	insertaddress := `INSERT INTO public.user_address(
						address_details, postal_code, country, user_id, province_id)
						VALUES ($1, $2, $3, $4, $5);`
	_, err = r.DB.Exec(insertaddress, addAddress.AddressDetails, addAddress.PostalCode, addAddress.Country, user_id, province_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

//TODO : DELETE, EDIT ACCOUNT AND USER
