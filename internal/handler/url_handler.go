package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/api"
	"go-short-url/internal/service"
	"go-short-url/util"
	"strconv"
)

type URLHandler struct {
	URLService service.URLService
}

func NewURLHandler(urlService service.URLService) *URLHandler {
	return &URLHandler{URLService: urlService}
}

func (URLHandler *URLHandler) CreateShortURL(ctx *fiber.Ctx) error {
	var request api.URLInputRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, api.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	shortURL, err := URLHandler.URLService.Create(context.Background(), request.LongURL)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, api.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := util.ResponseFormat(fiber.StatusCreated, api.MsgCreated, shortURL)
	return ctx.JSON(response)
}

func (URLHandler *URLHandler) RedirectURL(ctx *fiber.Ctx) error {
	shortCode := ctx.Params("shortCode")
	shortURL, err := URLHandler.URLService.GetLongURL(context.Background(), shortCode)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Redirect(shortURL.LongURL, fiber.StatusTemporaryRedirect)
}

func (URLHandler *URLHandler) GetById(ctx *fiber.Ctx) error {
	idParam := ctx.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, api.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := URLHandler.URLService.GetById(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusNotFound, api.MsgNotFound, nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}
	urlResponse := api.URLResponse{
		LongURL:   result.LongURL,
		ShortURL:  result.ShortCode,
		CreatedAt: result.CreatedAt,
	}
	response := util.ResponseFormat(fiber.StatusOK, api.MsgOk, urlResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (URLHandler *URLHandler) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, api.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	_, err = URLHandler.URLService.GetById(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusNotFound, api.MsgNotFound, nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}
	err = URLHandler.URLService.Delete(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, api.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := util.ResponseFormat(fiber.StatusOK, api.MsgOk, nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
