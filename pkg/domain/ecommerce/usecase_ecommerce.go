package ecommerce

import (
	"fmt"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
)

const TOPIC = "topic.ecommerce"

type EcommerceUseCase struct {
	repo  Repository
	event event.Event
}

func NewUseCase(repo Repository, event event.Event) *EcommerceUseCase {
	return &EcommerceUseCase{
		repo:  repo,
		event: event,
	}
}

func (usecase *EcommerceUseCase) Create(e *entity.Ecommerce) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	if err = usecase.repo.Create(e); err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err := usecase.event.Publish(TOPIC, e); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
