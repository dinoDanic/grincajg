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
)

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

	log.Fatal(app.Listen(":8000"))
}
