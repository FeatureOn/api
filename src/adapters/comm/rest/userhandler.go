package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/gorilla/mux"
)

// swagger:route GET /user/{id} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// GetUser handles GET requests
func (ctx *APIContext) GetUser(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	userService := application.NewUserService(ctx.userRepo)
	user, err := userService.GetUser(id)
	if err == nil {
		ToJSON(user, rw)
	}
}

// swagger:route PUT /user/{id} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// AddUser creates a new user on the system
func (ctx *APIContext) AddUser(rw http.ResponseWriter, r *http.Request) {
	// Get user data from oayload
	user, err := extractUserPayload(r)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	userService := application.NewUserService(ctx.userRepo)
	err = userService.AddUser(*user)
	if err == nil {
		ToJSON(user, rw)
	}
}

// extractConsentPayload extracts user data from the request body
// Returns user model if found, error otherwise
func extractUserPayload(r *http.Request) (user *domain.User, e error) {
	payload, e := readPayload(r)
	if e != nil {
		return
	}
	err := json.Unmarshal(payload, &user)
	if err != nil {
		e = errors.New(viper.GetString("CannotParsePayloadMsg"))
		log.Error().Err(err).Msg(viper.GetString("CannotParsePayloadMsg"))
		return
	}
	return
}
