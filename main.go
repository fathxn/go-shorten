package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-short-url/internal/config"
	"go-short-url/internal/handler"
	"go-short-url/internal/middleware"
	"go-short-url/internal/repository"
	"go-short-url/internal/service"
)

func main() {
	config.InitConfig()
	db := config.InitDB()

	urlRepository := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authHandler := handler.NewAuthHandler(userService)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${ip} ${status} - ${method} ${path}\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// user auth routes
	v1.Post("/auth/register", userHandler.RegisterUser)
	v1.Post("/auth/login", authHandler.AuthLogin)

	// short url routes
	v1.Post("/short_url", middleware.AuthMiddleware, urlHandler.CreateShortURL)
	v1.Get("/:user_id", middleware.AuthMiddleware, urlHandler.GetByUserId) // get list of shortened url by user_id
	v1.Delete("/short_url/:id", middleware.AuthMiddleware, urlHandler.Delete)
	app.Get("/:shortCode", urlHandler.RedirectURL) // redirect

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	_ = app.Listen(":1232")
}
