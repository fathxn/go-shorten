package middleware

import (
	"go-shorten/internal/domain"
	"go-shorten/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

	claimsMap := claims.Claims.(jwt.MapClaims)

	userId := claimsMap["id"].(string)
	ctx.Locals("userId", userId)

	// ctx.Locals("claims", claims)
	return ctx.Next()
}

func RequireVerifiedEmail(userUsecase domain.UserUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)
		user, err := userUsecase.GetById(c.Context(), userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get user"})
		}

		if !user.IsVerified {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "email not verified"})
		}

		return c.Next()
	}
}
