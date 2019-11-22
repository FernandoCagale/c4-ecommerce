package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-ecommerce/api/render"
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
	"github.com/FernandoCagale/c4-ecommerce/pkg/domain/ecommerce"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"net/http"
)

type EcommerceHandler struct {
	usecase ecommerce.UseCase
}

func NewEcommerce(usecase ecommerce.UseCase) *EcommerceHandler {
	return &EcommerceHandler{
		usecase: usecase,
	}
}

func (handler *EcommerceHandler) Ecommerces(w http.ResponseWriter, r *http.Request) {
	render.Response(w, map[string]bool{"ok": true}, http.StatusOK)
}

func (handler *EcommerceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ecommerce *entity.Ecommerce

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ecommerce); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(ecommerce); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, ecommerce, http.StatusCreated)
}
