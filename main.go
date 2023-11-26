package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-short-url/internal/config"
	"go-short-url/internal/handler"
	"go-short-url/internal/repository"
	"go-short-url/internal/service"
)

func main() {
	config.InitConfig()
	db := config.InitDB()

	urlRepository := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${ip} ${status} - ${method} ${path}\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/shorturl", urlHandler.CreateShortURL)
	v1.Get("/shorturl", urlHandler.GetById)
	v1.Delete("/shorturl/:id", urlHandler.Delete)
	app.Get("/:shortCode", urlHandler.RedirectURL)

	_ = app.Listen(":1234")
}
