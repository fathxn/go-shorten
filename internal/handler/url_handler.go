package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-short-url/internal/model/api"
	"go-short-url/internal/service"
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
		return err
	}

	shortURL, err := URLHandler.URLService.Create(context.Background(), request.LongURL)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response := api.APIResponse{
		Code:    fiber.StatusCreated,
		Message: "created",
		Data: api.URLResponse{
			LongURL:  request.LongURL,
			ShortURL: shortURL.ShortCode,
		},
	}

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
		return ctx.Status(fiber.StatusBadRequest).JSON(api.APIResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid parameter",
			Data:    struct{}{},
		})
	}

	result, err := URLHandler.URLService.GetById(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(api.APIResponse{
			Code:    fiber.StatusNotFound,
			Message: fmt.Sprintf("no record found with id %v", id),
			Data:    struct{}{},
		})
	}

	response := &api.APIResponse{
		Code:    fiber.StatusOK,
		Message: "ok",
		Data: api.URLResponse{
			LongURL:   result.LongURL,
			ShortURL:  result.ShortCode,
			CreatedAt: result.CreatedAt,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (URLHandler *URLHandler) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(api.APIResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "error",
			Data:    struct{}{},
		})
	}

	err = URLHandler.URLService.Delete(context.Background(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(api.APIResponse{
			Code:    fiber.StatusNotFound,
			Message: fmt.Sprintf("no record deleted with id %v", id),
			Data:    struct{}{},
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(api.APIResponse{
		Code:    fiber.StatusOK,
		Message: "ok",
		Data:    struct{}{},
	})
}
