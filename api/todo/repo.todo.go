package todo

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type TodoRepo struct {
	*sqlx.DB
}

func (r *TodoRepo) CreateCard(body *Card) (int64, error) {
	q0 := `INSERT INTO public.card (name) VALUES(:name)`
	res, err := r.NamedExec(q0, body)
	if err != nil {
		return 1, fiber.NewError(402, fmt.Sprintf("fail create card: %s", err.Error()))
	}

	return res.RowsAffected()
}
