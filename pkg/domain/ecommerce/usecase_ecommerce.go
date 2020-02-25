package ecommerce

import (
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
)

const EXCHANGE = "ecommerce"

type EcommerceUseCase struct {
	event event.Event
}

func NewUseCase(event event.Event) *EcommerceUseCase {
	return &EcommerceUseCase{
		event: event,
	}
}

func (usecase *EcommerceUseCase) Create(e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	if err := usecase.event.PublishExchange(EXCHANGE, e); err != nil {
		return err
	}

	return nil
}
