package customer

import (
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
)

const QUEUE = "notify.customer"

type CustomerUseCase struct {
	event event.Event
}

func NewUseCase(event event.Event) *CustomerUseCase {
	return &CustomerUseCase{
		event: event,
	}
}

func (usecase *CustomerUseCase) Create(e *entity.Customer) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	if err := usecase.event.PublishQueue(QUEUE, e); err != nil {
		return err
	}

	return nil
}
