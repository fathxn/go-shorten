package http

import (
	"context"
	"go-short-url/internal/model/dto"
	"go-short-url/internal/service"
	"go-short-url/util"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{UserService: userService}
}

func (h *userHandler) GetURLsByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")
	url, err := h.UserService.GetURLsByUserId(context.Background(), userId)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, dto.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	var urlResponse []dto.URLResponseByUserId
	for _, u := range *url {
		urlResponse = append(urlResponse, dto.URLResponseByUserId{
			Id:        u.Id,
			UserId:    u.UserId,
			LongURL:   u.LongURL,
			ShortURL:  u.ShortCode,
			CreatedAt: u.CreatedAt,
		})
	}

	response := util.ResponseFormat(fiber.StatusOK, dto.MsgOk, urlResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
