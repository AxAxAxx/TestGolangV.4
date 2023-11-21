package controllers

import (
	"github.com/AxAxAxx/go-test-api/modules/users/usecases"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUsecase usecases.ProductUsecase
}

func NewProductHandler(ProductUsecase usecases.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

func (p *ProductHandler) GetProduct(c *fiber.Ctx) error {
	productID := c.Params("productID")
	gender := c.Params("gender")
	style := c.Params("style")
	size := c.Params("size")
	limit := c.Params("limit")
	products, err := p.ProductUsecase.GetProducts(productID, gender, style, size, limit)
	if err != nil {
		return err
	}
	return c.JSON(products)
}
