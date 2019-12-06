package customer

import (
	"fmt"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
)

const TOPIC = "topic.customer"

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

	if err := usecase.event.Publish(TOPIC, e); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
