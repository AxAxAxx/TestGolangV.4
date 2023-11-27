package repositories

import (
	"fmt"
	"log"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoty struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepositoty {
	return &ProductRepositoty{
		DB: db,
	}
}

func (r *ProductRepositoty) GetProducts(id, gender, style, size, limit string, product []entities.Product) ([]entities.Product, error) {
	query := `SELECT p.product_id, sg.group_name, s.style_name, ps.size, g.gender_name, p.price, p.stock
	FROM public.product p JOIN style_product s ON p.styleproduct_id = s.style_id 
	JOIN gender_product g ON p.gender_id = g.gender_id 
	JOIN product_size ps ON p.productsize_id = ps.size_id 
	JOIN style_group sg ON s.stylegroup_id = sg.stylegroup_id WHERE 1=1`
	if id != "" {
		query += fmt.Sprintf(" AND p.product_id = '%s'", id)
	}
	if style != "" {
		query += fmt.Sprintf(" AND s.style_name = '%s'", style)
	}
	if gender != "" {
		query += fmt.Sprintf(" AND g.gender_name = '%s'", gender)
	}
	if size != "" {
		query += fmt.Sprintf(" AND ps.size = '%s'", size)
	}
	if limit != "" {
		query += fmt.Sprintf(" LIMIT '%s'", limit)
	}
	err := r.DB.Select(&product, query)
	if err != nil {
		log.Fatal(err)
	}
	return product, nil
}

func (r *ProductRepositoty) CreateProduct(newProduct entities.Product) error {
	_, err := r.DB.Exec(`INSERT INTO public.product(price, styleproduct_id, productsize_id, gender_id)
			VALUES ($1, $2, $3, $4);`, newProduct.Price, newProduct.StyleID, newProduct.SizeID, newProduct.GenderID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoty) UpdateProduct(productID int, updatedProduct entities.Product) error {
	_, err := r.DB.Exec(`UPDATE public.product SET price= $1, styleproduct_id= $2, productsize_id= $3, gender_id= $4, stock=$5 WHERE product_id = $6;`,
		updatedProduct.Price, updatedProduct.StyleID, updatedProduct.SizeID, updatedProduct.GenderID, updatedProduct.Stock, productID)
	if err != nil {
		return err
	}
	return err
}

func (r *ProductRepositoty) DeleteProduct(product_ID int) error {
	_, err := r.DB.Exec("DELETE FROM public.product WHERE product_id = $1;", product_ID)
	if err != nil {
		return err
	}
	return nil
}
