package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antoniohauren/finances/internal/services"
)

type Handlers struct {
	services *services.Services
}

func New(services *services.Services) *Handlers {
	return &Handlers{
		services: services,
	}
}

func (h Handlers) registerRoot() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello world")
	})
}

func (h Handlers) Listen(port int) error {
	h.registerRoot()

	h.registerUsersEndpoints()
	h.registerBillsEndpoints()
	h.registerPaymentsEndpoints()

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
