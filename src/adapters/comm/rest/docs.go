// Service classification of Toggler API
//
// Documentation for Toggler API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package rest

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	middleware "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/middleware"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response OK
type okResponseWrapper struct {
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body middleware.ValidationError
}

// A list of users
// swagger:response UsersResponse
type usersResponseWrapper struct {
	// All current products
	// in: body
	Body []dto.UserResponse
}

// Data structure representing a single user
// swagger:response UserResponse
type userResponseWrapper struct {
	// Newly created user
	// in: body
	Body dto.UserResponse
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateUser
type userIDParamsWrapper struct {
	// The id of the user for which the operation relates
	// in: path
	// required: true
	ID string `json:"id"`
}
