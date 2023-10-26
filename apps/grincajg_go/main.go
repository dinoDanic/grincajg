package main

import (
	"fmt"
	"grincajg/controllers"
	"grincajg/database"
	"grincajg/middleware"
	"log"

	_ "grincajg/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Grincajg API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	// database.DB.AutoMigrate(&models.User{}, &models.Organization{}, &models.Store{})
}

func serveApplication() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
	})

	micro.Route("/organization", func(router fiber.Router) {
		router.Post("/", middleware.DeserializeUser, controllers.CreateOrganization)
		router.Get("/", middleware.DeserializeUser, controllers.GetOrganization)
		router.Post("/create-store", middleware.DeserializeUser, controllers.CreateStore)
		router.Get("/stores", middleware.DeserializeUser, controllers.GetOrganizationStores)
	})

	micro.Route("/users", func(router fiber.Router) {
		router.Get("/me", middleware.DeserializeUser, controllers.GetMe)
	})

	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":8080"))
}
