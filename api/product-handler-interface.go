package api

import (
	"strconv"

	"github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductBody struct {
	Name			string		`json:"name" xml:"name" form:"name"`
	Price			float64		`json:"price" xml:"price" form:"price"`
	Quantity		int			`json:"quantity" xml:"quantity" form:"quantity"`
	Section_code	string		`json:"section_code" xml:"section_code" form:"section_code"`
}

type ProductHandler interface {
	Get(*fiber.Ctx) error
	Post(*fiber.Ctx) error
	Put(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	GetAll(*fiber.Ctx) error
}

type producthandler struct {
	productService domain.ProductService
}

func NewProductHandler(productService domain.ProductService) ProductHandler {
	return &producthandler{productService: productService}
}

// Get func get a product.
// @Description Get a product.
// @Summary get a product
// @Tags Product
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} domain.Product
// @Security ApiKeyAuth
// @Router /v1/product/{email} [get]
func (h *producthandler) Get(c *fiber.Ctx) error {
	email := c.Params("email")
	result, err := h.productService.Find(email)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid retrieving a product details",
		})
	}
	return c.JSON(result)
}

// Post func create a product.
// @Description Create a product.
// @Summary create a product
// @Tags Product
// @Accept json
// @Produce json
// @Success 201
// @Security ApiKeyAuth
// @Router /v1/product [post]
func (h *producthandler) Post(c *fiber.Ctx) error {
	pb := new(ProductBody)
    if err := c.BodyParser(pb); err != nil {
        return c.JSON(Response{
			Code: 400,
			Description: err.Error(),
		})
    }

	product, _ := h.productService.Find(pb.Code)
	if product.Code != "" {
		return c.JSON(Response{
			Code: 400,
			Description: "Unsuccessful creating of product. Product might already exist.",
		})
	}

	err := h.productService.Store(pb.Code, pb.Name, pb.Price, pb.Quantity, pb.Section_code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: err.Error(),
		})
	}
	return c.JSON(Response{
		Code: 201,
		Description: "Successfully create a product",
	})
}

// Put func update a product.
// @Description Update a product.
// @Summary update a product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/product [put]
func (h *producthandler) Put(c *fiber.Ctx) error {
	pb := new(ProductBody)
    if err := c.BodyParser(pb); err != nil {
        return c.JSON(Response{
			Code: 400,
			Description: err.Error(),
		})
    }
	
	err := h.productService.Update(pb.Code, pb.Name, pb.Price, pb.Quantity, pb.Section_code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid update product",
		})
	}
	return c.JSON(Response{
		Code: 200,
		Description: "Successfully update a product",
	})
}

// Delete func remove a product.
// @Description Remove a product.
// @Summary remove a product
// @Tags Product
// @Accept json
// @Produce json
// @Param code path string true "code"
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/product/{code} [delete]
func (h *producthandler) Delete(c *fiber.Ctx) error {
	code := c.Params("code")
	err := h.productService.Delete(code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid delete product",
		})
	}
	return c.JSON(Response{
		Code: 200,
		Description: "Successfully delete a product",
	})
}

// Get func gets all products.
// @Description Gets all products.
// @Summary gets all products
// @Tags Product
// @Accept json
// @Produce json
// @Param section_code path string true "Section Code"
// @Param page path int true "Page"
// @Param page_size path int true "Page Size"
// @Success 200 {array} domain.Product
// @Security ApiKeyAuth
// @Router /v1/product/section/{section_code}/{page}/{page_size} [get]
func (h *producthandler) GetAll(c *fiber.Ctx) error {
	product_code := c.Params("product_code")
	page, _ := strconv.Atoi(c.Params("page"))
	page_size, _ := strconv.Atoi(c.Params("page_size"))
	results, err := h.productService.FindAll(product_code, page, page_size)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid get all products by section",
		})
	}
	return c.JSON(results)
}