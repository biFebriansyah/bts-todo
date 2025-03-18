package users

import (
	"fmt"

	"github.com/biFebriansyah/bts-todoapp/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	*UserRepo
}

func (r *UserHandler) SignIn(ctx *fiber.Ctx) error {
	body := new(User)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userData, err := r.GetByUsename(body.Usename)
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(body.Password, userData.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "worng password")
	}

	token, err := utils.GenerateJwt(userData.UserId)
	if err != nil {
		return fmt.Errorf("fail when generate token: %w", err)
	}

	return ctx.JSON(fiber.Map{"token": token})
}

func (r *UserHandler) SignUp(ctx *fiber.Ctx) error {
	body := new(User)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	pass, err := utils.HashPassword(body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body.Password = pass
	userData, err := r.CreateUser(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"users": userData})
}
