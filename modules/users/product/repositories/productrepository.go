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

func (r *ProductRepositoty) GetProducts(filter []string, sorting, limit, offset string, product []entities.ProductRes) ([]entities.ProductRes, error) {
	query := `SELECT 
				p.product_id, 
				sg.group_name, 
				s.style_name, 
				ps.size, 
				g.gender_name, 
				p.price, 
				p.stock
				FROM public.product p 
				JOIN style_product s ON p.styleproduct_id = s.style_id
				JOIN gender_product g ON p.gender_id = g.gender_id
				JOIN product_size ps ON p.productsize_id = ps.size_id
				JOIN style_group sg ON s.stylegroup_id = sg.stylegroup_id 
				WHERE 1=1 `

	q := []string{"AND p.product_id =", "AND s.style_name =", "AND g.gender_name =", "AND ps.size ="}
	var (
		queryCondition, empty []string
		valueFilter           []interface{}
		temp, result          string
	)

	for index, value := range filter {
		if value != "" {
			valueFilter = append(valueFilter, value)
			for indexQuery, valueQuery := range q {
				if index == indexQuery {
					queryCondition = append(queryCondition, valueQuery)
				}
			}
		}
	}
	for index, value := range queryCondition {
		temp += func(i int) string {
			empty = append(empty, value)
			if i != len(queryCondition)-1 {
				if empty != nil {
					result = fmt.Sprintf("%s $%d ", empty[i], i+1)
				}
				return result
			} else {
				return fmt.Sprintf("%s $%d", empty[i], i+1)
			}
		}(index)
	}
	query += temp
	if sorting != "" {
		query += fmt.Sprintf(" ORDER BY p.product_id %s", sorting)
	}
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	if offset != "" {
		query += fmt.Sprintf(" OFFSET %s", offset)
	}
	rows, err := r.DB.Queryx(query, valueFilter...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p entities.ProductRes
		err := rows.StructScan(&p)
		if err != nil {
			return nil, err
		}
		product = append(product, p)
	}
	return product, nil
}

func (r *ProductRepositoty) CreateProduct(newProduct entities.Product) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Println("Panic:", p)
		}
	}()
	_, err = tx.Exec(`INSERT INTO public.product(price, styleproduct_id, productsize_id, gender_id, stock)
			VALUES ($1, $2, $3, $4, $5);`, newProduct.Price, newProduct.StyleID, newProduct.SizeID, newProduct.GenderID, newProduct.Stock)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil
	}
	return nil
}

func (r *ProductRepositoty) UpdateProduct(productID int, existingProduct entities.Product) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Println("Panic:", p)
		}
	}()
	_, err = r.DB.Exec(`UPDATE public.product
		SET price=$1, styleproduct_id=$2, productsize_id=$3, gender_id=$4, stock=$5
		WHERE product_id=$6;`,
		existingProduct.Price, existingProduct.StyleID, existingProduct.SizeID, existingProduct.GenderID, existingProduct.Stock, productID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *ProductRepositoty) ExistingProduct(productID int, existingProduct entities.Product) (entities.Product, error) {
	query := `SELECT price, styleproduct_id, productsize_id, gender_id, stock 
			FROM public.product 
			WHERE product_id = $1`
	err := r.DB.Get(&existingProduct, query, productID)
	if err != nil {
		return existingProduct, err
	}
	return existingProduct, nil
}

func (r *ProductRepositoty) DeleteProduct(product_ID int) error {
	_, err := r.DB.Exec("DELETE FROM public.product WHERE product_id = $1;", product_ID)
	if err != nil {
		return err
	}
	return nil
}
