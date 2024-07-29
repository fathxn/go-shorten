package http

import (
	"context"
	"go-shorten/internal/model/dto"
	"go-shorten/internal/service"
	"go-shorten/util"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{AuthService: authService}
}

func (h *authHandler) RegisterUser(ctx *fiber.Ctx) error {
	request := &dto.UserRegisterInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := util.ErrorValidation(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := h.AuthService.RegisterUser(context.Background(), request); err != nil {
		response := util.ResponseFormat(fiber.StatusConflict, "email already registered", err.Error())
		return ctx.Status(fiber.StatusConflict).JSON(response)
	}

	response := util.ResponseFormat(fiber.StatusCreated, dto.MsgCreated, nil)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *authHandler) AuthLogin(ctx *fiber.Ctx) error {
	request := &dto.UserLoginInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := util.ErrorValidation(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	user, err := h.AuthService.LoginUser(context.Background(), request)
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
