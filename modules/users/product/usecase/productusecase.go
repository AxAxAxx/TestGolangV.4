package usecase

import (
	"strconv"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/product/repositories"
)

type ProductUsecase struct {
	ProductRepository repositories.ProductRepositoty
}

func NewProductUsecase(productRepository repositories.ProductRepositoty) *ProductUsecase {
	return &ProductUsecase{
		ProductRepository: productRepository,
	}
}

func (u *ProductUsecase) GetProducts(filter []string, sorting, limit, offset string) ([]entities.ProductRes, error) {
	products, err := u.ProductRepository.GetProducts(filter, sorting, limit, offset, []entities.ProductRes{})
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (u *ProductUsecase) CreateProduct(newProduct entities.Product) error {
	err := u.ProductRepository.CreateProduct(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) UpdateProduct(productID int, existingProduct, updatedProduct entities.Product) error {
	product, err := u.ProductRepository.ExistingProduct(productID, existingProduct)
	if err != nil {
		return err
	}

	empty, _ := strconv.Atoi("")
	if updatedProduct.Price != empty {
		product.Price = updatedProduct.Price
	}
	if updatedProduct.StyleID != empty {
		product.StyleID = updatedProduct.StyleID
	}
	if updatedProduct.SizeID != empty {
		product.SizeID = updatedProduct.SizeID
	}
	if updatedProduct.GenderID != empty {
		product.GenderID = updatedProduct.GenderID
	}
	if updatedProduct.Stock != empty {
		product.Stock = updatedProduct.Stock
	}
	err = u.ProductRepository.UpdateProduct(productID, product)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductUsecase) DeleteProduct(product_ID int) error {
	err := u.ProductRepository.DeleteProduct(product_ID)
	if err != nil {
		return err
	}
	return nil
}
