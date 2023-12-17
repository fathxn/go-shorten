package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/dto"
	"go-short-url/internal/service"
	"go-short-url/util"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (AuthHandler *AuthHandler) RegisterUser(ctx *fiber.Ctx) error {
	request := &dto.UserRegisterInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := AuthHandler.AuthService.RegisterUser(context.Background(), request); err != nil {
		response := util.ResponseFormat(fiber.StatusConflict, "email already registered", nil)
		return ctx.Status(fiber.StatusConflict).JSON(response)
	}

	response := util.ResponseFormat(fiber.StatusCreated, dto.MsgCreated, nil)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (AuthHandler *AuthHandler) AuthLogin(ctx *fiber.Ctx) error {
	request := &dto.UserLoginInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	user, err := AuthHandler.AuthService.LoginUser(context.Background(), request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusUnauthorized, dto.MsgUnauthorized, nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	token, err := util.GenerateJWT(user.Id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	authResponse := dto.UserAuthResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	response := util.ResponseFormat(fiber.StatusOK, dto.MsgOk, authResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
