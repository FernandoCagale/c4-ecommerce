package ecommerce

import (
	"fmt"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
)

const TOPIC = "topic.ecommerce"

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

	if err := usecase.event.Publish(TOPIC, e); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
