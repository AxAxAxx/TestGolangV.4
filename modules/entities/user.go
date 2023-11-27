package entities

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
