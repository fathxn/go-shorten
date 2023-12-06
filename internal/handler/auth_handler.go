package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/api"
	"go-short-url/internal/service"
	"go-short-url/util"
)

type AuthHandler struct {
	UserService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

func (AuthHandler *AuthHandler) AuthLogin(ctx *fiber.Ctx) error {
	request := &api.UserLoginInput{}
	err := ctx.BodyParser(request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, api.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	user, err := AuthHandler.UserService.LoginUser(context.Background(), request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusUnauthorized, api.MsgUnauthorized, nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}
	token, err := util.GenerateJWT(user.Id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, api.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	authResponse := api.UserAuthResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	response := util.ResponseFormat(fiber.StatusOK, api.MsgOk, authResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
