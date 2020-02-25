package entity

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Ecommerce struct {
	Customer CustomerEcommerce `json:"customer"`
	Order    Order             `json:"order"`
}

func (e Ecommerce) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Customer, validation.Required),
		validation.Field(&e.Order),
	)
}

type CustomerEcommerce struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (e CustomerEcommerce) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Name, validation.Required, validation.Length(1, 50)),
	)
}

type Order struct {
	Code  string    `json:"code"`
	Date  time.Time `json:"date"`
	Items []*Item   `json:"items"`
}

func (e Order) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Date, validation.Required),
		validation.Field(&e.Items, validation.Required, validation.Each()),
	)
}

type Item struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Value       float64 `json:"value"`
}

func (e Item) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Code, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Description, validation.Required, validation.Length(1, 50)),
		validation.Field(&e.Quantity, validation.Required),
		validation.Field(&e.Value, validation.Required),
	)
}
