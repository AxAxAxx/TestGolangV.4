package controllers

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/product/repositories"
)

type ProductController struct {
	ProductRepository repositories.ProductRepositoty
}

func NewProductController(ProductRepository repositories.ProductRepositoty) *ProductController {
	return &ProductController{
		ProductRepository: ProductRepository,
	}
}

func (u *ProductController) GetProducts(id, gender, style, size, limit string) ([]entities.Product, error) {
	products, err := u.ProductRepository.GetProducts(id, gender, style, size, limit, []entities.Product{})
	if err != nil {
		return nil, err
	}
	return products, nil
}

// func (u *ProductController) GetProductByID(productID string) ([]entities.Product, error) {
// 	products, err := u.ProductRepository.GetProductByID(productID, []entities.Product{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, nil
// }

func (u *ProductController) CreateProduct(newProduct entities.Product) error {
	err := u.ProductRepository.CreateProduct(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductController) UpdateProduct(productID int, updatedProduct entities.Product) error {
	err := u.ProductRepository.UpdateProduct(productID, updatedProduct)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductController) DeleteProduct(product_ID int) error {
	err := u.ProductRepository.DeleteProduct(product_ID)
	if err != nil {
		return err
	}
	return nil
}
