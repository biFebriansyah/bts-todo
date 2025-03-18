package todo

import "github.com/gofiber/fiber/v2"

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
