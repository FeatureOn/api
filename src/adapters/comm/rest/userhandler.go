package rest

import (
	"net/http"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"
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
