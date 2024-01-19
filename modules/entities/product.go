package entities

type Product struct {
	ProductID int `db:"product_id"`
	Price     int `db:"price"`
	StyleID   int `db:"styleproduct_id"`
	SizeID    int `db:"productsize_id"`
	GenderID  int `db:"gender_id"`
	Stock     int `db:"stock"`
}

type ProductRes struct {
	ProductID    int    `json:"product_id" db:"product_id"`
	StyleGroup   string `json:"group_name" db:"group_name"`
	StyleProduct string `json:"style_name" db:"style_name"`
	Size         string `json:"size" db:"size"`
	Gender       string `json:"gender_name" db:"gender_name"`
	Price        int    `json:"price" db:"price"`
	Stock        int    `json:"stock" db:"stock"`
}
