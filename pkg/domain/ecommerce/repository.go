package ecommerce

import "github.com/FernandoCagale/c4-ecommerce/pkg/entity"

type Repository interface {
	Create(ecommerce *entity.Ecommerce) (err error)
}
