package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-short-url/util"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		response := util.ResponseFormat(fiber.StatusUnauthorized, "unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	claims, err := util.VerifyJWT(tokenString)
	if err != nil {
		response := util.ResponseFormat(fiber.StatusUnauthorized, "unauthorized", nil)
		return ctx.Status(fiber.StatusUnauthorized).JSON(response)
	}

	ctx.Locals("claims", claims)
	return ctx.Next()
}
