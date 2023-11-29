package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-short-url/internal/config"
	"go-short-url/internal/handler"
	api2 "go-short-url/internal/model/api"
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

	userLogin := api2.UserLoginInput{
		Email:    "syrotul.inayah@gmail.com",
		Password: "kanaya",
	}
	user, err := userService.LoginUser(context.Background(), &userLogin)
	if err != nil {
		return
	}
	fmt.Println(user.Name)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${ip} ${status} - ${method} ${path}\n",
	}))
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// user auth routes
	v1.Post("/auth/register", userHandler.RegisterUser)
	v1.Post("/auth/login", userHandler.LoginUser)

	// short url routes
	v1.Post("/short_url", urlHandler.CreateShortURL)
	v1.Get("/short_url", urlHandler.GetById)
	v1.Delete("/short_url/:id", urlHandler.Delete)
	app.Get("/:shortCode", urlHandler.RedirectURL)

	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	_ = app.Listen(":1232")
}
