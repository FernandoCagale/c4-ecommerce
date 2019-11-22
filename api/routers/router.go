package routers

import (
	"github.com/FernandoCagale/c4-ecommerce/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler *handlers.HealthHandler
	ecommerceHandler *handlers.EcommerceHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/ecommerce", routes.ecommerceHandler.Create).Methods("POST")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, ecommerceHandler *handlers.EcommerceHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler: healthHandler,
		ecommerceHandler: ecommerceHandler,
	}
}
