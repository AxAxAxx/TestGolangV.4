package usecases

import (
	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/product/repositories"
)

type ProductUsecase struct {
	ProductRepository repositories.ProductRepositoty
}

func NewProductUsecase(ProductRepository repositories.ProductRepositoty) *ProductUsecase {
	return &ProductUsecase{
		ProductRepository: ProductRepository,
	}
}

func (u *ProductUsecase) GetProducts(productID, gender, style, size, limit string) ([]entities.Product, error) {
	products, err := u.ProductRepository.GetProducts(productID, gender, style, size, limit, []entities.Product{})
	if err != nil {
		return nil, err
	}
	return products, nil
}
