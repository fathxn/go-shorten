package http

import (
	"context"
	"go-shorten/internal/domain"
	"go-shorten/internal/domain/errors"
	"go-shorten/util"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	UserService domain.UserUsecase
}

func NewUserHandler(userService domain.UserUsecase) *userHandler {
	return &userHandler{UserService: userService}
}

func (h *userHandler) RegisterUser(ctx *fiber.Ctx) error {
	request := &domain.UserRegisterInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, errors.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := util.ErrorValidation(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, errors.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if _, err := h.UserService.RegisterUser(context.Background(), request); err != nil {
		response := util.ResponseFormat(fiber.StatusConflict, "email already registered", err.Error())
		return ctx.Status(fiber.StatusConflict).JSON(response)
	}

	response := util.ResponseFormat(fiber.StatusCreated, errors.MsgCreated, nil)
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *userHandler) AuthLogin(ctx *fiber.Ctx) error {
	request := &domain.UserLoginInput{}
	if err := ctx.BodyParser(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, errors.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if err := util.ErrorValidation(request); err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, errors.MsgInternalServerError, err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	user, err := h.UserService.LoginUser(context.Background(), request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusUnauthorized, errors.MsgUnauthorized, nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	token, err := util.GenerateJWT(user.Id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, errors.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	authResponse := domain.UserAuthResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	response := util.ResponseFormat(fiber.StatusOK, errors.MsgOk, authResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *userHandler) GetURLsByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")
	url, err := h.UserService.GetURLsByUserId(context.Background(), userId)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, errors.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	var urlResponse []domain.URLResponseByUserId
	for _, u := range *url {
		urlResponse = append(urlResponse, domain.URLResponseByUserId{
			Id:        u.Id,
			UserId:    u.UserId,
			LongURL:   u.LongURL,
			ShortURL:  u.ShortCode,
			CreatedAt: u.CreatedAt,
		})
	}

	response := util.ResponseFormat(fiber.StatusOK, errors.MsgOk, urlResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
