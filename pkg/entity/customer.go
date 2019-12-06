package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Customer struct {
	Code   string   `json:"code"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Phone  string   `json:"phone"`
	Notify []string `json:"notify"`
}

func (e Customer) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Email, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Phone, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Notify, validation.Required),
	)
}
