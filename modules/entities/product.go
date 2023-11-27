package entities

type Product struct {
	ProductID    int    `db:"product_id"`
	StyleGroup   string `db:"group_name"`
	StyleProduct string `db:"style_name"`
	Size         string `db:"size"`
	Gender       string `db:"gender_name"`
	Price        int    `db:"price"`
	StyleID      int    `db:"styleproduct_id"`
	SizeID       int    `db:"productsize_id"`
	GenderID     int    `db:"gender_id"`
	Stock        int    `db:"stock"`
}
