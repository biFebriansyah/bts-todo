package users

import "time"

type User struct {
	UserId    string     `json:"user_id"`
	Usename   string     `json:"usename"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Users []User
