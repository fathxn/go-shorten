package main

import (
	"go-shorten/config"
	"go-shorten/internal/delivery/http"
	"go-shorten/internal/middleware"
	"go-shorten/internal/repository/postgres"
	"go-shorten/internal/usecase"
	"go-shorten/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// config.InitConfig()
	// db := config.InitDB()
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize repositories
	urlRepository := postgres.NewURLRepository(db)
	userRepository := postgres.NewUserRepository(db)

	// Initialize usecases
	urlUsecase := usecase.NewURLUsecase(urlRepository)
	userUsecase := usecase.NewUserUsecase(userRepository, urlRepository)
	authUsecase := usecase.NewAuthUsecase(userRepository)

	// Initialize handlers
	urlHandler := http.NewURLHandler(urlUsecase)
	userHandler := http.NewUserHandler(userUsecase)
	authHandler := http.NewAuthHandler(authUsecase)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format: "${ip} ${status} - ${method} ${path}\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// auth routes
	v1.Post("/auth/register", authHandler.RegisterUser)
	v1.Post("/auth/login", authHandler.AuthLogin)

	// user routes
	v1.Get("/:user_id", userHandler.GetURLsByUserId) // get list of shortened url by user_id

	// short url routes
	v1.Post("/short_url", middleware.AuthMiddleware, urlHandler.CreateShortURL)
	v1.Get("/short_url/", urlHandler.GetById) // short_url/id?=
	v1.Delete("/short_url/:id", urlHandler.Delete)
	app.Get("/:shortCode", urlHandler.RedirectURL) // redirect

	_ = app.Listen("127.0.0.1:1232")
}
