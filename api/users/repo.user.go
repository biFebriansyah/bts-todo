package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	*sqlx.DB
}

func (r *UserRepo) CreateUser(body *User) (int64, error) {
	q0 := `INSERT INTO public.users (username, email, "password") VALUES(:username, :email, :password);`
	res, err := r.NamedExec(q0, body)
	if err != nil {
		return 1, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("fail create card: %s", err.Error()))
	}

	return res.RowsAffected()
}

func (r *UserRepo) GetByUsename(username string) (*User, error) {
	q0 := `SELECT user_id, username, email, "password", created_at, updated_at FROM public.users where username=$1`

	data := new(User)
	if err := r.Get(data, q0, username); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("fail get cards: %s", err.Error()))
		}
		return nil, fiber.NewError(fiber.StatusBadGateway, fmt.Sprintf("fail get cards: %s", err.Error()))
	}

	return data, nil
}
