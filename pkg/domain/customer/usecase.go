package customer

import (
	"github.com/FernandoCagale/c4-ecommerce/pkg/entity"
)

type UseCase interface {
	Create(customer *entity.Customer) (err error)
}
