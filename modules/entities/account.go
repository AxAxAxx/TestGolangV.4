package entities

import "time"

type User struct {
	UserID         int    `db:"user_id"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	Email          string `db:"email"`
	PhoneNumber    string `db:"phonenumber"`
	DateOfBirth    string `db:"dete_of_birth"`
	AddressDetails string `db:"address_details"`
	PostalCode     string `db:"postal_code"`
	Province       string `db:"province"`
	Country        string `db:"country"`
}

type RegisterAccount struct {
	AccountID      int       `db:"account_id"`
	Username       string    `db:"username"`
	Password       string    `db:"password"`
	CratedAt       time.Time `db:"crated_at"`
	UserID         int       `db:"user_id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	DateOfBirth    string    `db:"dete_of_birth"`
	PhoneNumber    string    `db:"phonenumber"`
	AddressDetails string    `db:"address_details"`
	PostalCode     string    `db:"postal_code"`
	Province       string    `db:"province_name"`
	ProvinceID     int       `db:"province_id"`
	Country        string    `db:"country"`
}
