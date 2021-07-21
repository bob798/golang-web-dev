package user

import "github.com/go-playground/validator/v10"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,mail"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
