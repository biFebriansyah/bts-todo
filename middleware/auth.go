package middleware

import (
	"strings"

	"github.com/biFebriansyah/bts-todoapp/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()["Authorization"]
	if len(headers) <= 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "need login")
	}

	tokens := strings.Replace(headers[0], "Bearer ", "", 1)
	claims, err := utils.ParseJwt(tokens)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	ctx.Locals("userId", claims.Id)
	return ctx.Next()
}
