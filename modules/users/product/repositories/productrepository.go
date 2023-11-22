package repositories

import (
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

func (r *ProductRepositoty) GetProducts(productID, gender, style, size, limit string, product []entities.Product) ([]entities.Product, error) {
	query := `SELECT p.product_id, sg.group_name, s.style_name, ps.size, g.gender_name, p.price 
	FROM public.product p JOIN style_product s ON p.styleproduct_id = s.style_id 
	JOIN gender_product g ON p.gender_id = g.gender_id 
	JOIN product_size ps ON p.productsize_id = ps.size_id 
	JOIN style_group sg ON s.stylegroup_id = sg.stylegroup_id`
	rows, err := r.DB.Queryx(query)
	if style != "" && size != "" && gender != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND ps.size = $2 AND g.gender_name = $3 ORDER BY p.product_id LIMIT $4`, style, size, gender, limit)
	} else if style != "" && size != "" && gender == "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND ps.size = $2 ORDER BY p.product_id LIMIT $3`, style, size, limit)
	} else if style != "" && gender != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND g.gender_name = $2 ORDER BY p.product_id LIMIT $3`, style, gender, limit)
	} else if size != "" && gender != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE ps.size = $1 AND g.gender_name = $2 ORDER BY p.product_id LIMIT $3`, size, gender, limit)
	} else if style != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 ORDER BY p.product_id LIMIT $2`, style, limit)
	} else if size != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE ps.size = $1 ORDER BY p.product_id LIMIT $2`, size, limit)
	} else if gender != "" && limit != "" {
		rows, err = r.DB.Queryx(query+` WHERE g.gender_name = $1  ORDER BY p.product_id LIMIT $2`, gender, limit)
	} else if style != "" && size != "" && gender != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND ps.size = $2 AND g.gender_name = $3 ORDER BY p.product_id`, style, size, gender)
	} else if style != "" && size != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND ps.size = $2 ORDER BY p.product_id`, style, size)
	} else if style != "" && gender != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 AND g.gender_name = $2 ORDER BY p.product_id`, style, gender)
	} else if style == "" && size != "" && gender != "" {
		rows, err = r.DB.Queryx(query+` WHERE ps.size = $1 AND g.gender_name = $2 ORDER BY p.product_id`, size, gender)
	} else if style != "" {
		rows, err = r.DB.Queryx(query+` WHERE s.style_name = $1 ORDER BY p.product_id`, style)
	} else if size != "" {
		rows, err = r.DB.Queryx(query+` WHERE ps.size = $1 ORDER BY p.product_id`, size)
	} else if gender != "" {
		rows, err = r.DB.Queryx(query+` WHERE g.gender_name = $1  ORDER BY p.product_id`, gender)
	} else if productID != "" {
		rows, err = r.DB.Queryx(query+` WHERE p.product_id = $1 ORDER BY p.product_id`, productID)
	}
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p entities.Product
		err := rows.Scan(&p.ProductID, &p.StyleGroup, &p.StyleProduct, &p.Size, &p.Gender, &p.Price)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		product = append(product, p)
	}
	return product, nil
}
