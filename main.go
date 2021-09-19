package main

import (
	config "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/pkg/configs"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/pkg/middleware"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/pkg/utils"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/repository"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/domain"
    "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/api"
	_ "github.com/leonardyeoxl/CRUD-app-designed-with-Hexagonal-Architecture-in-Golang/docs"
	
	swagger "github.com/arsmn/fiber-swagger/v2"
    "github.com/gofiber/fiber/v2"
    _ "github.com/joho/godotenv/autoload"                // load .env file automatically
)

// @title Grocery Shop Management System API
// @version 1.0
// @description System to mange Sections and Products.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email leonardyeoxl@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	sectionRepo, productRepo,_ := repository.NewPostgreSQLRepository(config.DatabaseConfig())
	sectionService := domain.NewSectionService(sectionRepo)
	sectionHandler := api.NewSectionHandler(sectionService)
    productService := domain.NewProductService(productRepo)
	productHandler := api.NewProductHandler(productService)

	// Define Fiber config.
    fiber_config := config.FiberConfig()

    // Define a new Fiber app with config.
    app := fiber.New(fiber_config)

    // Middlewares.
    middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

    // Create routes group
	route := app.Group("api/v1")

	// Create routes group.
    swagger_route := app.Group("/swagger")

    // Routes for GET method:
    swagger_route.Get("*", swagger.Handler) // get one user by ID

    // Route for access token
    route.Get("/token/new", api.GetNewAccessToken) // create a new access tokens

	// Routes for Section
	route.Get("/section/:code", middleware.JWTProtected(), sectionHandler.Get) // retrieve a section
	route.Post("/section/:name", middleware.JWTProtected(), sectionHandler.Post) // create a new section
	route.Delete("/section/:code", middleware.JWTProtected(), sectionHandler.Delete) //delete a existing section
	route.Get("/section", middleware.JWTProtected(), sectionHandler.GetAll) // retrieve all sections
	route.Put("/section/:code/:name", middleware.JWTProtected(), sectionHandler.Put) // update a existing section
	
	// Routes for Product
	route.Get("/product/:email", middleware.JWTProtected(), productHandler.Get) // retrieve a product
	route.Post("/product/:email/:name/:merchant_code", middleware.JWTProtected(), productHandler.Post) // create a new product of section
	route.Delete("/product/:email", middleware.JWTProtected(), productHandler.Delete) //delete a existing product
	route.Get("/product/merchant/:merchant_code/:page/:page_size", middleware.JWTProtected(), productHandler.GetAll) // retrieve all products of section
	route.Put("/product/:email/:name", middleware.JWTProtected(), productHandler.Put) // update a existing product

    // Start server.
    utils.StartServer(app)
}