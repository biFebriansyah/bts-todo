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
	var userId string
	if userId = ctx.Locals("userId").(string); userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please login")
	}

	body := new(Card)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body.UserId = userId
	result, err := h.CreateCard(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) GetCards(ctx *fiber.Ctx) error {
	var userId string
	if userId = ctx.Locals("userId").(string); userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please login")
	}
	result, err := h.GetCard(userId)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) DeleteCards(ctx *fiber.Ctx) error {
	var userId string
	if userId = ctx.Locals("userId").(string); userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "please login")
	}

	uid := fiberutil.CopyString(ctx.Params("uid"))
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

	cardId := fiberutil.CopyString(ctx.Params("cardId"))
	body.CardId = cardId
	result, err := h.CreateTodo(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) GetTodo(ctx *fiber.Ctx) error {
	cardId := fiberutil.CopyString(ctx.Params("cardId"))
	result, err := h.GetAllTodos(cardId)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) GetTodoId(ctx *fiber.Ctx) error {
	cardId := fiberutil.CopyString(ctx.Params("cardId"))
	itemId := fiberutil.CopyString(ctx.Params("itemId"))
	result, err := h.GetTodosById(cardId, itemId)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) UpdateStatus(ctx *fiber.Ctx) error {
	uid := fiberutil.CopyString(ctx.Params("itemId"))

	body := new(Todo)
	body.TodoId = uid
	body.Status = true
	result, err := h.UpdateTodoStatus(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) UpdateName(ctx *fiber.Ctx) error {
	body := new(Todo)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	uid := fiberutil.CopyString(ctx.Params("itemId"))
	body.TodoId = uid
	result, err := h.UpdateTodoName(body)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}

func (h *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	uid := fiberutil.CopyString(ctx.Params("itemId"))
	result, err := h.DeleteTodoItem(uid)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"result": result})
}
