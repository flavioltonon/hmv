package analysts

import (
	"net/http"

	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/drivers"
	"flavioltonon/hmv/infrastructure/middleware"

	"github.com/gorilla/mux"
)

type Controller struct {
	usecases *Usecases
	drivers  *drivers.Drivers
}

type Usecases struct {
	Authentication usecases.AuthenticationUsecase
	Analysts       usecases.AnalystUsecase
	Emergencies    usecases.EmergencyUsecase
}

func NewController(usecases *Usecases, drivers *drivers.Drivers) *Controller {
	return &Controller{usecases: usecases, drivers: drivers}
}

func (c *Controller) SetRoutes(parent *mux.Router) {
	parent.Use(mux.MiddlewareFunc(middleware.Authentication(
		c.usecases.Authentication,
		c.drivers.Logger,
		c.drivers.Presenter,
	)))

	parent.HandleFunc("/emergencies", c.listEmergencies).Methods(http.MethodGet)
	parent.HandleFunc("/emergencies", c.listEmergencies).Methods(http.MethodGet)
}