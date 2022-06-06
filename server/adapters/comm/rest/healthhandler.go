package rest

import (
	"net/http"

	"github.com/FeatureOn/api/server/application"

	"github.com/rs/zerolog/log"
)

// swagger:route GET /health/live Health Live
// Return 200 if the api is up and running
// responses:
//	200: OK
//	404: errorResponse

// Live handles GET requests
func (apiContext *APIContext) Live(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

// swagger:route GET /health/ready Health Ready
// Return 200 if the api is up and running and connected to the database
// responses:
//	200: OK
//	404: errorResponse

// Ready handles GET requests
func (apiContext *APIContext) Ready(rw http.ResponseWriter, r *http.Request) {
	hs := application.NewHealthService(apiContext.healthRepo)
	status := hs.Ready()
	if status == false {
		log.Error().Msg("Error connecting to database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
