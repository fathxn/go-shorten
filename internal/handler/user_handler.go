package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/api"
	"go-short-url/internal/service"
	"go-short-url/util"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (UserHandler *UserHandler) RegisterUser(ctx *fiber.Ctx) error {
	userInput := &api.UserRegisterInput{}
	if err := ctx.BodyParser(userInput); err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, api.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := UserHandler.UserService.RegisterUser(context.Background(), userInput)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusConflict, "email already registered", nil)
		return ctx.Status(fiber.StatusConflict).JSON(response)
	}

	response := util.ResponseFormat(fiber.StatusCreated, api.MsgCreated, nil)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}
