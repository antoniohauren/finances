package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antoniohauren/finances/internal/services"
)

type Controller struct {
	services *services.Services
}

func New(services *services.Services) *Controller {
	return &Controller{
		services: services,
	}
}

func (c Controller) registerRoot() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello world")
	})
}

func (c Controller) Listen(port int) error {
	c.registerRoot()

	c.registerUsersEndpoints()

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
