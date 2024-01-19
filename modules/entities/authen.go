package entities

import "time"

type LoginRequest struct {
	Username string `db:"username" form:"username"`
	Password string `db:"password" form:"password"`
}

type AccountRequest struct {
	Username  string `db:"username"`
	Password  string `db:"password"`
	AccountID int    `db:"account_id"`
	UserID    int    `db:"user_id"`
	Role_id   int    `db:"role_id"`
}

type Account struct {
	Id       int       `db:"account_id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	CratedAt time.Time `db:"crated_at"`
	UserID   int       `db:"user_id"`
	RoleID   int       `db:"role_id"`
}

type Token struct {
	AccountID    int    `db:"account_id"`
	UserID       int    `db:"user_id"`
	Role_id      int    `db:"role_id"`
	Username     string `db:"username"`
	Token        string `db:"token"`
	RefreshToken string `db:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `db:"refresh_token"`
}

type DeleteToken struct {
	Token        string `db:"token"`
	RefreshToken string `db:"refresh_token"`
}
