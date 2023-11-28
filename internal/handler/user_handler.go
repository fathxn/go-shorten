package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/api"
	"go-short-url/internal/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (UserHandler *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	userInput := &api.UserInput{}
	if err := ctx.BodyParser(userInput); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(api.APIResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}

	err := UserHandler.UserService.Create(context.Background(), userInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(api.APIResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    struct{}{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(api.APIResponse{
		Code:    fiber.StatusOK,
		Message: "created",
		Data:    struct{}{},
	})
}
