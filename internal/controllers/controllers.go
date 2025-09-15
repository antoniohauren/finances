package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c Controller) registerRoot() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello world")
	})
}

func (c Controller) Listen(port int) error {
	c.registerRoot()

	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
