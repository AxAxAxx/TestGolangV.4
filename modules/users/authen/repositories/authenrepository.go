package repositories

import (
	"log"

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

func (r *AccountRepositoty) FindByUsername(username entities.LoginRequest) (entities.AccountRequest, error) {
	var user entities.AccountRequest
	err := r.DB.Get(&user, `SELECT 
								ac.account_id, 
								ac.username, 
								ac.password, 
								ac.role_id, 
								u.user_id 
							FROM account ac 
							JOIN "user" u ON u.account_id = ac.account_id 
							WHERE username = $1`, username.Username)
	if err != nil {
		// return user, err
		log.Fatal(err)
	}
	return user, nil
}

func (r *AccountRepositoty) Token(id int, access_token, refresh_token string) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	_, err = r.DB.Exec(`INSERT INTO tokens (account_id, token, refresh_token) VALUES ($1, $2, $3)`, id, access_token, refresh_token)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

func (r *AccountRepositoty) FindRefreshToken(refresh_token entities.RefreshTokenRequest) (entities.Token, error) {
	var token entities.Token
	err := r.DB.Get(&token, `SELECT tk.account_id, u.user_id, ac.username, tk.token, tk.refresh_token, ac.role_id
								FROM public.tokens tk
								LEFT JOIN account ac ON ac.account_id = tk.account_id
								JOIN "user" u ON u.account_id = ac.account_id 
								WHERE tk.refresh_token = $1`, refresh_token.RefreshToken)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (r *AccountRepositoty) UpdateToken(access_token, refresh_token string, id entities.Token) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	_, err = r.DB.Exec("UPDATE tokens SET token=$1, refresh_token=$2 WHERE account_id = $3", access_token, refresh_token, id.AccountID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

func (r *AccountRepositoty) DeleteToken(access_token, refresh_token string) error {
	_, err := r.DB.Exec("DELETE FROM public.tokens WHERE token = $1 AND refresh_token = $2;", access_token, refresh_token)
	if err != nil {
		return err
	}
	return nil
}
