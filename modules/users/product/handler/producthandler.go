package handler

import (
	"fmt"
	"strconv"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/product/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(ProductUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	gender := c.Query("gender")
	style := c.Query("style")
	size := c.Query("size")
	limit := c.Query("limit")
	offset := c.Query("offset")
	sorting := c.Query("sorting")
	filter := []string{id, gender, style, size}
	products, err := h.ProductUsecase.GetProducts(filter, sorting, limit, offset)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	return c.JSON(entities.ProductsSuccessResponse(&products))
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var newProduct entities.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	err := h.ProductUsecase.CreateProduct(newProduct)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	fmt.Println(newProduct)
	return c.JSON(entities.ProductSuccessResponse(newProduct))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	var _, updatedProduct entities.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}
	// err = h.ProductUsecase.UpdateProduct(productID, existingProduct, updatedProduct)
	// if err != nil {
	// 	return c.JSON(entities.ErrorResponse(err))
	// }

	return c.JSON(updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString("Invalid Product ID")
		return c.JSON(entities.ErrorResponse(err))
	}

	err = h.ProductUsecase.DeleteProduct(bookID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		return c.JSON(entities.ErrorResponse(err))
	}

	return c.JSON(entities.DeleteResponse())
}
