package todo

import "time"

type Card struct {
	CardId    string     `db:"card_id" json:"card_id,omitempty"`
	UserId    string     `db:"user_id" json:"user_id,omitempty"`
	CardName  string     `db:"card_name" json:"card_name"`
	TodoList  []Todo     `db:"todo_list" json:"todo_list" xml:"todo_list" form:"todo_list"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Cards []Card

type Todo struct {
	TodoId    string     `db:"todo_id" json:"todo_id,omitempty"`
	CardId    string     `db:"card_id" json:"card_id,omitempty"`
	TodoName  string     `db:"todo_name" json:"todo_name"`
	Status    bool       `db:"todo_status" json:"todo_status"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Todos []Todo
