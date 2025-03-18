package todo

import "time"

type Card struct {
	CardId    string     `json:"card_id,omitempty"`
	CardName  string     `json:"card_name"`
	TodoList  []Todo     `json:"todo_list" xml:"todo_list" form:"todo_list"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Cards []Card

type Todo struct {
	TodoId    string     `json:"todo_id,omitempty"`
	CardId    string     `json:"card_id,omitempty"`
	TodoName  string     `json:"todo_name"`
	Done      bool       `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
