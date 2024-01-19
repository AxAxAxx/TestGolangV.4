package repositories

import (
	"fmt"
	"time"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/jmoiron/sqlx"
)

type OrderRepositoty struct {
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepositoty {
	return &OrderRepositoty{
		DB: db,
	}
}

// Create Order
func (r *OrderRepositoty) CreateOrders(user_id float64, created_at, expired time.Time, shipping, newOrder entities.Order) (int, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return 0, nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	err = r.DB.QueryRowx(`INSERT INTO public."order"(
							user_id, 
							shipping_details, 
							order_status_id, 
							created_at, 
							expired)
						VALUES ($1, $2, $3, $4, $5) RETURNING order_id;`,
		user_id, shipping.ShippingDetails, newOrder.OrderStatusID, created_at, expired).Scan(&newOrder.OrderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, nil
	}
	return newOrder.OrderID, nil
}

func (r *OrderRepositoty) GetShipDetails(user_id float64, shipdetails entities.Order) (entities.Order, error) {
	q := `SELECT jsonb_agg(address)
	FROM (SELECT ua.address_details, ua.postal_code, pv.province_name, ua.country
	FROM "user" u
	Join user_address ua ON ua.user_id = u.user_id
	Join province pv ON ua.province_id = pv.province_id
	WHERE u.user_id = $1) as address`

	err := r.DB.Get(&shipdetails.ShippingDetails, q, user_id)
	if err != nil {
		return shipdetails, err
	}
	return shipdetails, nil
}

func (r *OrderRepositoty) AddProduct(order_id, total_price int, t_price []int, created_at time.Time, order_product entities.Order) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	for indexOrder, productReq := range order_product.Products {
		for indexTotalprice, price := range t_price {
			if indexTotalprice == indexOrder {
				total_price = price
			}
		}
		_, err := r.DB.Exec(`INSERT INTO order_product 
							(order_id, product_id, total_price, quantity, created_at) 
							VALUES ($1, $2, $3, $4, $5)`,
			order_id, productReq.ProductID, total_price, productReq.Quantity, created_at)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

func (r *OrderRepositoty) GetPriceAndStock(ps entities.Price_Stock, op entities.Order) ([]entities.Price_Stock, error) {
	var price_stock []entities.Price_Stock
	for _, productReq := range op.Products {
		queryproduct := `SELECT price, stock FROM public.product WHERE product_id = $1;`
		err := r.DB.Get(&ps, queryproduct, productReq.ProductID)
		if err != nil {
			return nil, err
		}
		price_stock = append(price_stock, ps)
	}
	return price_stock, nil
}

func (r *OrderRepositoty) UpdatedStock(updatedStock []int, op entities.Order) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	for index, productReq := range op.Products {
		for i, value := range updatedStock {
			if i == index {
				_, err := r.DB.Exec(`UPDATE public.product SET stock=$1 WHERE product_id =$2`, value, productReq.ProductID)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

// Get OrderByAdmin
func (r *OrderRepositoty) GetOrdersByAdmid(filter []string, limit, sorting, offset string, order []entities.OrderRes) ([]entities.OrderRes, error) {

	var (
		queryCondition, empty []string
		valueFilter           []interface{}
		temp, s               string
	)

	q := []string{"AND o.order_id =", "AND u.first_name =", "AND u.last_name =", "AND os.status =",
		"AND u.phonenumber =", "AND created_at =", "AND expired ="}

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
					s = fmt.Sprintf("%s $%d ", empty[i], i+1)
				}
				return s
			} else {
				return fmt.Sprintf("%s $%d", empty[i], i+1)
			}
		}(index)
	}

	query := `SELECT
    			o.order_id,
				o.user_id,
				u.first_name,
				u.last_name,
				u.phonenumber,
				o.shipping_details,
				o.order_status_id,
				os.status,
				o.created_at,
				o.expired,
				(
					SELECT
						jsonb_agg(x)
					FROM (
							SELECT
								op.product_id,
								p.product_id,
								s.style_name,
								ps.size,
								g.gender_name,
								op.quantity,
								op.total_price
							FROM product p
								LEFT JOIN order_product op ON op.product_id = p.product_id
								LEFT JOIN style_product s ON p.styleproduct_id = s.style_id
								LEFT JOIN gender_product g ON p.gender_id = g.gender_id
								LEFT JOIN product_size ps ON p.productsize_id = ps.size_id
							WHERE op.order_id = o.order_id
						) AS x
				) AS product_details
			FROM "order" o
				LEFT JOIN "user" u ON o.user_id = u.user_id
				LEFT JOIN order_status os ON o.order_status_id = os.orderstatus_id
			WHERE 1=1 `

	query += temp
	if sorting != "" {
		query += fmt.Sprintf(" ORDER BY o.order_id %s", sorting)
	}
	if limit != "" {
		query += fmt.Sprintf(" LIMIT %s", limit)
	}
	if offset != "" {
		query += fmt.Sprintf(" OFFSET %s", offset)
	}
	fmt.Println(query)
	rows, err := r.DB.Queryx(query, valueFilter...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var o entities.OrderRes
		err := rows.StructScan(&o)
		if err != nil {
			return nil, err
		}
		order = append(order, o)
	}
	return order, nil
}

// Get Order By user
func (r *OrderRepositoty) GetOrders(id float64, order []entities.OrderRes) ([]entities.OrderRes, error) {
	query := `SELECT
    			o.order_id,
				o.user_id,
				u.first_name,
				u.last_name,
				u.phonenumber,
				o.shipping_details,
				o.order_status_id,
				os.status,
				o.created_at,
				o.expired,
				(
					SELECT
						jsonb_agg(x)
					FROM (
							SELECT
								op.product_id,
								p.product_id,
								s.style_name,
								ps.size,
								g.gender_name,
								op.quantity,
								op.total_price
							FROM product p
								LEFT JOIN order_product op ON op.product_id = p.product_id
								LEFT JOIN style_product s ON p.styleproduct_id = s.style_id
								LEFT JOIN gender_product g ON p.gender_id = g.gender_id
								LEFT JOIN product_size ps ON p.productsize_id = ps.size_id
							WHERE op.order_id = o.order_id
						) AS x
				) AS product_details
			FROM "order" o
				LEFT JOIN "user" u ON o.user_id = u.user_id
				LEFT JOIN order_status os ON o.order_status_id = os.orderstatus_id
			WHERE u.user_id = $1 `
	rows, err := r.DB.Queryx(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var o entities.OrderRes
		err := rows.StructScan(&o)
		if err != nil {
			return nil, err
		}
		order = append(order, o)
	}
	return order, nil
}

//TODO : DELETE Order
