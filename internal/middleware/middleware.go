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
