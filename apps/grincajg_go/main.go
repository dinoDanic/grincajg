package main

import (
	"fmt"
	"grincajg/controllers"
	"grincajg/database"
	"grincajg/env"
	"grincajg/middleware"
	"grincajg/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "grincajg/docs"
)

// @title Grincajg API 2
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api
func main() {
	config := loadEnv()
	loadDatabase(config)
	serveApplication()
}

func loadEnv() env.Config {
	config, err := env.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	return config
}

func loadDatabase(config env.Config) {
	database.Connect(config)
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Organization{})
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

	micro.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)
	micro.Post("/organization", middleware.DeserializeUser, controllers.CreateOrganization)

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

	log.Fatal(app.Listen(":8000"))
}
