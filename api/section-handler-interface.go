package api

import (
	"strings"
	
	"github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Response struct {
	Code		uint
	Description	string
}

type SectionHandler interface {
	Get(*fiber.Ctx) error
	Post(*fiber.Ctx) error
	Put(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	GetAll(*fiber.Ctx) error
}

type sectionhandler struct {
	sectionService domain.SectionService
}

func NewSectionHandler(sectionService domain.SectionService) SectionHandler {
	return &sectionhandler{sectionService: sectionService}
}

// Get func get a section.
// @Description Get a section.
// @Summary get a section
// @Tags Section
// @Accept json
// @Produce json
// @Param code path string true "Code"
// @Success 200 {object} domain.Section
// @Security ApiKeyAuth
// @Router /v1/section/{code} [get]
func (h *sectionhandler) Get(c *fiber.Ctx) error {
	code := c.Params("code")
	result, err := h.sectionService.Find(code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid retrieving a section details",
		})
	}
	return c.JSON(result)
}

// Post func create a section.
// @Description Create a section.
// @Summary create a section
// @Tags Section
// @Accept json
// @Produce json
// @Param name path string true "Name"
// @Success 201
// @Security ApiKeyAuth
// @Router /v1/section/{name} [post]
func (h *sectionhandler) Post(c *fiber.Ctx) error {
	name := c.Params("name")
	uuidWithHyphen := uuid.New()
    code := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	err := h.sectionService.Store(name, code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid creating section",
		})
	}
	return c.JSON(Response{
		Code: 201,
		Description: "Successfully create a section",
	})
}

// Put func update a section.
// @Description Update a section.
// @Summary update a section
// @Tags Section
// @Accept json
// @Produce json
// @Param code path string true "Code"
// @Param name path string true "Name"
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/section/{code}/{name} [put]
func (h *sectionhandler) Put(c *fiber.Ctx) error {
	code := c.Params("code")
	name := c.Params("name")
	err := h.sectionService.Update(code, name)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid update section",
		})
	}
	return c.JSON(Response{
		Code: 200,
		Description: "Successfully update a section",
	})
}

// Delete func remove a section.
// @Description Remove a section.
// @Summary remove a section
// @Tags Section
// @Accept json
// @Produce json
// @Param code path string true "Code"
// @Success 200
// @Security ApiKeyAuth
// @Router /v1/section/{code} [delete]
func (h *sectionhandler) Delete(c *fiber.Ctx) error {
	code := c.Params("code")
	err := h.sectionService.Delete(code)
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid delete section",
		})
	}
	return c.JSON(Response{
		Code: 200,
		Description: "Successfully delete a section",
	})
}

// Get func gets all sections.
// @Description Gets all sections.
// @Summary gets all sections
// @Tags Section
// @Accept json
// @Produce json
// @Success 200 {array} domain.Section
// @Security ApiKeyAuth
// @Router /v1/section/ [get]
func (h *sectionhandler) GetAll(c *fiber.Ctx) error {
	results, err := h.sectionService.FindAll()
	if err != nil {
		return c.JSON(Response{
			Code: 400,
			Description: "Invalid delete section",
		})
	}
	return c.JSON(results)
}