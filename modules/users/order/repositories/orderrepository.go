package repositories

import (
	"encoding/json"
	"log"
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

func (r *OrderRepositoty) CreateOrder(newOrder entities.Order) error {
	newOrder.CreatedAt = time.Now()
	newOrder.StartDate = time.Now()
	newOrder.EndDate = newOrder.StartDate.Add(24 * time.Hour)
	var productdetails entities.Product
	var user entities.User

	queryproduct := `SELECT s.style_name, ps.size, g.gender_name, p.price
	FROM public.product p JOIN style_product s ON p.styleproduct_id = s.style_id
	JOIN gender_product g ON p.gender_id = g.gender_id
	JOIN product_size ps ON p.productsize_id = ps.size_id
	WHERE p.product_id = $1`

	err := r.DB.QueryRowx(queryproduct, newOrder.ProductID).Scan(&productdetails.StyleProduct, &productdetails.Size, &productdetails.Gender, &productdetails.Price)
	if err != nil {
		return err
	}

	newOrder.Total_Price = newOrder.Quantity * productdetails.Price

	productdetail := entities.ProductDetails{
		StyleProduct: productdetails.StyleProduct,
		Size:         productdetails.Size,
		Gender:       productdetails.Gender,
		Price:        productdetails.Price,
	}

	jsonDataProduct, err := json.Marshal(productdetail)
	if err != nil {
		log.Fatal(err)
	}

	queryshipaddress := `SELECT u.first_name, u.last_name, u.phonenumber, ua.address_details, ua.postal_code, ua.province, ua.country
	FROM "user" u Join user_address ua ON ua.user_id = u.user_id
	WHERE u.user_id = $1`

	err = r.DB.QueryRowx(queryshipaddress, newOrder.UserID).Scan(&user.FirstName, &user.LastName, &user.PhoneNumber, &user.AddressDetails, &user.PostalCode, &user.Province, &user.Country)
	if err != nil {
		return err
	}

	shipdetails := entities.ShippingDetails{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		AddressDetails: user.AddressDetails,
		PostalCode:     user.PostalCode,
		Province:       user.Province,
		Country:        user.Country,
	}

	jsonDataShip, err := json.Marshal(shipdetails)
	if err != nil {
		log.Fatal(err)
	}

	err = r.DB.Get(&productdetails, "SELECT stock FROM product WHERE product_id = $1", newOrder.ProductID)
	if err != nil {
		return err
	}

	if productdetails.Stock < newOrder.Quantity {
		return err
	}

	updatedStock := productdetails.Stock - newOrder.Quantity
	_, err = r.DB.Exec("UPDATE product SET stock = $1 WHERE product_id = $2", updatedStock, newOrder.ProductID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(`INSERT INTO "order" (product_details, shipping_details, created_at, user_id, product_id, quantity, total_price, orderstatus_id, start_date, end_date) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		jsonDataProduct, jsonDataShip, newOrder.CreatedAt, newOrder.UserID, newOrder.ProductID, newOrder.Quantity, newOrder.Total_Price, newOrder.OrderStatusID, newOrder.StartDate, newOrder.EndDate)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepositoty) GetOrders(id string, fname string, lname string, phonenumber string, status string, startdate string, enddate string, limit string, order []entities.Order) ([]entities.Order, error) {
	q := `SELECT o.order_id, o.user_id, u.first_name, u.last_name, u.phonenumber, o.shipping_details, o.product_id,o.orderstatus_id, o.product_details, o.created_at, o.quantity, o.total_price, os.status, o.start_date, o.end_date
		FROM public."order" o Join "user" u ON o.user_id = u.user_id
		JOIN order_status os ON o.orderstatus_id = os.orderstatus_id`
	rows, err := r.DB.Queryx(q)
	if id != "" {
		rows, err = r.DB.Queryx(q+` WHERE o.order_id = $1`, id)
	} else if fname != "" && limit != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.first_name = $1 ORDER BY o.order_id LIMIT $2`, fname, limit)
	} else if lname != "" && limit != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.last_name = $1 ORDER BY o.order_id LIMIT $2`, lname, limit)
	} else if phonenumber != "" && limit != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.phonenumber = $1 ORDER BY o.order_id LIMIT $2`, phonenumber, limit)
	} else if startdate != "" && enddate != "" && limit != "" {
		rows, err = r.DB.Queryx(q+` WHERE o.start_date = $1 AND o.end_date = $2 ORDER BY o.order_id LIMIT $2`, startdate, enddate, limit)
	} else if status != "" && limit != "" {
		rows, err = r.DB.Queryx(q+` WHERE os.status = $1 ORDER BY o.order_id LIMIT $2`, status, limit)
	} else if fname != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.first_name = $1 ORDER BY o.order_id`, fname)
	} else if lname != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.last_name = $1 ORDER BY o.order_id`, lname)
	} else if phonenumber != "" {
		rows, err = r.DB.Queryx(q+` WHERE u.phonenumber = $1 ORDER BY o.order_id`, phonenumber)
	} else if startdate != "" && enddate != "" {
		rows, err = r.DB.Queryx(q+` WHERE o.start_date = $1 AND o.end_date = $2 ORDER BY o.order_id`, startdate, enddate)
	} else if status != "" {
		rows, err = r.DB.Queryx(q+` WHERE os.status = $1 ORDER BY o.order_id `, status)
	}
	if err != nil {
		log.Fatal("Failed to execute the query:", err)
	}
	defer rows.Close()

	var retrievedProduct []byte
	var retrievedSipping []byte

	for rows.Next() {
		var o entities.Order
		err := rows.Scan(&o.OrderID, &o.UserID, &o.FirstName, &o.LastName, &o.PhoneNumber, &retrievedSipping, &o.ProductID, &o.OrderStatusID, &retrievedProduct, &o.CreatedAt, &o.Quantity, &o.Total_Price, &o.OrderStatus, &o.StartDate, &o.EndDate)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		err = json.Unmarshal(retrievedProduct, &o.ProductDetails)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON data:", err)
		}
		err = json.Unmarshal(retrievedSipping, &o.ShippingDetails)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON data:", err)
		}
		order = append(order, o)
	}
	return order, nil
}
