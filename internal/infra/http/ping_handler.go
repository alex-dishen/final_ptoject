package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PingHandler
// @Summary  This endpoint can be used as health check for this application.
// @Tags     Health
// @Accept   json
// @Produce  json
// @Success  200  {string}  response
// @Router   /ping [get]
func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("OK")
		if err != nil {
			fmt.Printf("writing response: %s", err)
		}
	}
}
