package rest

import (
	"context"
	"net/http"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/middleware"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type ValidatedUser struct{}

// swagger:route GET /user/{id} User GetUser
// Return the user if found
// responses:
//	200: OK
//	404: errorResponse

// GetUser gets a single user if found
func (ctx *APIContext) GetUser(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	userService := application.NewUserService(ctx.userRepo)
	user, err := userService.GetUser(id)
	if err == nil {
		respondWithJSON(rw, r, 200, dto.MapUser2UserResponse(user))
	}
}

// swagger:route POST /user/{id} User AddUser
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddUser creates a new user on the system
func (ctx *APIContext) AddUser(rw http.ResponseWriter, r *http.Request) {
	// Get user data from oayload
	userDTO := r.Context().Value(ValidatedUser{}).(dto.UserRequest)
	user := dto.MapUserRequest2User(userDTO)
	userService := application.NewUserService(ctx.userRepo)
	err := userService.AddUser(user)
	if err == nil {
		respondWithJSON(rw, r, 200, user)
	}
}

// MiddlewareValidateNewUser Checks the integrity of new user in the request and calls next if ok
func (ctx *APIContext) MiddlewareValidateNewUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := middleware.ExtractUserPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the product
		errs := ctx.validation.Validate(user)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the user")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), ValidatedUser{}, *user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
