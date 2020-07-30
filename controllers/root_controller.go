package controllers

import (
	"net/http"

	"github.com/dennypradipta/cariliquid/responses"
)

// Root for Healthcheck
func (server *Server) Root(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API is Healthy")

}
