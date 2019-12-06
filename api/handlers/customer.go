package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-ecommerce/api/render"
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
	"github.com/FernandoCagale/c4-ecommerce/pkg/domain/customer"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"net/http"
)

type CustomerHandler struct {
	usecase customer.UseCase
}

func NewCustomer(usecase customer.UseCase) *CustomerHandler {
	return &CustomerHandler{
		usecase: usecase,
	}
}

func (handler *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var customer *entity.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(customer); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, customer, http.StatusCreated)
}
