package handlers

import (
	"github.com/hellofresh/health-go"
	"os"

	healthRabbitmq "github.com/hellofresh/health-go/checks/rabbitmq"

	"net/http"
	"time"
)

type HealthHandler struct {
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	health.Register(health.Config{
		Name:    "rabbitmq",
		Timeout: time.Second * 5,
		Check: healthRabbitmq.New(healthRabbitmq.Config{
			DSN: os.Getenv("AMQP_URI"),
		}),
	})

	health.HandlerFunc(w, r)
}
