package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-short-url/util"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "unauthorized"})
	}
	_, err := util.VerifyJWT(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"msg": "unauthorized"})
	}

	return ctx.Next()
}

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
