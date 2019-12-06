package ecommerce

import "github.com/google/wire"

var Set = wire.NewSet(NewUseCase,
	wire.Bind(new(UseCase), new(*EcommerceUseCase)))
