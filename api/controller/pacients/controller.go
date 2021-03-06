package pacients

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
	Pacients       usecases.PacientUsecase
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

	parent.HandleFunc("", c.createPacient).Methods(http.MethodPost)
	parent.HandleFunc("/{pacient_id}", c.findPacient).Methods(http.MethodGet)
	parent.HandleFunc("/{pacient_id}/emergency-contacts", c.updateEmergencyContact).Methods(http.MethodPut)
	parent.HandleFunc("/{pacient_id}/health", c.updateHealthData).Methods(http.MethodPut)
	parent.HandleFunc("/{pacient_id}/location", c.updateLocationData).Methods(http.MethodPut)
}
