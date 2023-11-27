package handler

import (
	"strconv"

	"github.com/AxAxAxx/go-test-api/modules/entities"
	"github.com/AxAxAxx/go-test-api/modules/users/product/controllers"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	ProductUsecase controllers.ProductController
}

func NewProductHandler(ProductUsecase controllers.ProductController) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: ProductUsecase,
	}
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Query("id", "")
	gender := c.Query("gender", "")
	style := c.Query("style", "")
	size := c.Query("size", "")
	limit := c.Query("limit", "")
	products, err := h.ProductUsecase.GetProducts(id, gender, style, size, limit)
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
	return c.JSON(entities.ProductSuccessResponse(newProduct))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("ID"))
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	var updatedProduct entities.Product
	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	err = h.ProductUsecase.UpdateProduct(productID, updatedProduct)
	if err != nil {
		return c.JSON(entities.ErrorResponse(err))
	}

	return c.JSON(entities.UpdateResponse())
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
