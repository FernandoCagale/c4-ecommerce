package ecommerce

import (
	"github.com/FernandoCagale/c4-ecommerce/internal/errors"
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
	"github.com/ThreeDotsLabs/watermill"
)

type InMemoryRepository struct {
	m map[string]*entity.Ecommerce
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]*entity.Ecommerce{}}
}

func (repo *InMemoryRepository) Create(e *entity.Ecommerce) (err error) {
	if e.ID == "" {
		e.ID = watermill.NewUUID()
	}
	if repo.m[e.ID] != nil {
		return errors.ErrInvalidPayload
	}
	repo.m[e.ID] = e
	return nil
}
