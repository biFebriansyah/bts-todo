package todo

import (
	"github.com/gofiber/fiber/v2"
	fiberutil "github.com/gofiber/fiber/v2/utils"
)

type TodoHandler struct {
	*TodoRepo
}

func NewHandler(r *TodoRepo) *TodoHandler {
	return &TodoHandler{r}
}

func (h *TodoHandler) AddCard(ctx *fiber.Ctx) error {
	body := new(Card)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.CreateCard(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) GetCards(ctx *fiber.Ctx) error {
	var userId string
	if userId = ctx.Locals("image").(string); userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please login")
	}
	result, err := h.GetCard(userId)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) DeleteCards(ctx *fiber.Ctx) error {
	uid := fiberutil.CopyString(ctx.Params("uuid"))
	result, err := h.DeleteCard(uid)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) AddTodo(ctx *fiber.Ctx) error {
	body := new(Todo)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.CreateTodo(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) GetTodo(ctx *fiber.Ctx) error {
	var userId string
	if userId = ctx.Locals("image").(string); userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please login")
	}

	result, err := h.GetTodos(userId)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}
