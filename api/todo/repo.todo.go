package todo

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type TodoRepo struct {
	*sqlx.DB
}

func NewRepo(db *sqlx.DB) *TodoRepo {
	return &TodoRepo{db}
}

func (r *TodoRepo) CreateCard(body *Card) (int64, error) {
	q0 := `INSERT INTO public.cards (user_id, card_name) VALUES(:user_id, :card_name)`
	res, err := r.NamedExec(q0, body)
	if err != nil {
		return 1, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("fail create card: %s", err.Error()))
	}

	return res.RowsAffected()
}

func (r *TodoRepo) GetCard(userId string) (*Cards, error) {
	q := `SELECT 
		c.card_id, 
		c.user_id, 
		c.card_name,
		(SELECT 
			JSONB_AGG(JSONB_BUILD_OBJECT(
				'todo_id', t.todo_id,
				'todo_name', t.todo_name,
				'todo_status', t.todo_status
			))
			FROM public.todos t
			WHERE c.card_id = t.card_id 
		) AS todo_list, 
		c.created_at, 
		c.updated_at
		FROM public.cards c
		WHERE c.user_id = $1`

	var datas = new(Cards)
	if rows, err := r.Queryx(q, userId); err == nil {
		for rows.Next() {
			var data = new(Card)
			var todolisJson []byte
			err := rows.Scan(
				&data.CardId,
				&data.UserId,
				&data.CardName,
				&todolisJson,
				&data.CreatedAt,
				&data.UpdatedAt,
			)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			if len(todolisJson) > 0 {
				if err := json.Unmarshal(todolisJson, &data.TodoList); err != nil {
					return nil, fmt.Errorf("failed to unmarshal music_artis: %w", err)
				}
			}

			*datas = append(*datas, *data)
		}
	}

	return datas, nil
}

func (r *TodoRepo) DeleteCard(uid string) (int64, error) {
	q := `DELETE FROM public.cards WHERE card_id = $1;`
	res := r.MustExec(q, uid)
	return res.RowsAffected()
}

func (r *TodoRepo) CreateTodo(body *Todo) (int64, error) {
	q0 := `INSERT INTO public.todos(card_id, todo_name, todo_status) VALUES(:card_id, :todo_name, :todo_status);`
	res, err := r.NamedExec(q0, body)
	if err != nil {
		return 1, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("fail create todos: %s", err.Error()))
	}

	return res.RowsAffected()
}

func (r *TodoRepo) GetTodos(userId string) (*[]Todo, error) {
	q := `SELECT todo_id, card_id, todo_name, todo_status, created_at, updated_at FROM public.todos WHERE card_id = $1`

	var data = new([]Todo)
	if err := r.Get(data, q, userId); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("fail get cards: %s", err.Error()))
		}
		return nil, fiber.NewError(fiber.StatusBadGateway, fmt.Sprintf("fail get cards: %s", err.Error()))
	}

	return data, nil
}

func (r *TodoRepo) UpdateTodo(body *Todo) (int64, error) {
	q0 := `
	UPDATE public.todos SET 
		todo_name=COALESCE(NULLIF(:todo_name, ''), todo_name),
		todo_status=COALESCE(NULLIF(:todo_status, ''), todo_status),
		updated_at=now()
	WHERE todo_id = :todo_id;`

	res, err := r.NamedExec(q0, body)
	if err != nil {
		return 1, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("fail update todos: %s", err.Error()))
	}

	return res.RowsAffected()
}

func (r *TodoRepo) DeleteTodo(uid string) (int64, error) {
	q := `DELETE FROM public.todos WHERE todo_id = $1;`
	res := r.MustExec(q, uid)
	return res.RowsAffected()
}
