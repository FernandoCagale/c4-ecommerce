package pkg

import (
	"github.com/FernandoCagale/c4-ecommerce/api/handlers"
	"github.com/FernandoCagale/c4-ecommerce/api/routers"
	"github.com/FernandoCagale/c4-ecommerce/internal/event"
	"github.com/FernandoCagale/c4-ecommerce/pkg/domain/customer"
	"github.com/FernandoCagale/c4-ecommerce/pkg/domain/ecommerce"
	"github.com/google/wire"
)

var Container = wire.NewSet(ecommerce.Set, handlers.Set, routers.Set, event.Set, customer.Set)
