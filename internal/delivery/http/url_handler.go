package http

import (
	"context"
	"go-short-url/internal/model/dto"
	"go-short-url/internal/service"
	"go-short-url/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type urlHandler struct {
	URLService service.URLService
}

func NewURLHandler(urlService service.URLService) *urlHandler {
	return &urlHandler{URLService: urlService}
}

func (h *urlHandler) CreateShortURL(ctx *fiber.Ctx) error {
	var request dto.URLInputRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, dto.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	shortURL, err := h.URLService.Create(context.Background(), request.LongURL)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := util.ResponseFormat(fiber.StatusCreated, dto.MsgCreated, shortURL)
	return ctx.JSON(response)
}

func (h *urlHandler) RedirectURL(ctx *fiber.Ctx) error {
	shortCode := ctx.Params("shortCode")
	shortURL, err := h.URLService.GetLongURL(context.Background(), shortCode)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Redirect(shortURL.LongURL, fiber.StatusTemporaryRedirect)
}

func (h *urlHandler) GetById(ctx *fiber.Ctx) error {
	idParam := ctx.Query("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusBadRequest, dto.MsgBadRequest, nil)
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}
	result, err := h.URLService.GetById(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusNotFound, dto.MsgNotFound, nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}
	urlResponse := dto.URLResponse{
		LongURL:   result.LongURL,
		ShortURL:  result.ShortCode,
		CreatedAt: result.CreatedAt,
	}
	response := util.ResponseFormat(fiber.StatusOK, dto.MsgOk, urlResponse)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *urlHandler) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	_, err = h.URLService.GetById(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusNotFound, dto.MsgNotFound, nil)
		return ctx.Status(fiber.StatusNotFound).JSON(response)
	}
	err = h.URLService.Delete(context.Background(), id)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusInternalServerError, dto.MsgInternalServerError, nil)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response := util.ResponseFormat(fiber.StatusOK, dto.MsgOk, nil)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
