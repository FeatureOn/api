package rest

import (
	"context"
	"net/http"

	"github.com/FeatureOn/api/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/adapters/comm/rest/mappers"
	middleware "github.com/FeatureOn/api/adapters/comm/rest/middleware"
	"github.com/FeatureOn/api/application"
	"github.com/rs/zerolog/log"
)

type validatedEnvironment struct{}

// swagger:route POST /user/{id} User AddUser
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddEnvironment creates a new environment on the system
func (ctx *APIContext) AddEnvironment(rw http.ResponseWriter, r *http.Request) {
	// Get environment data from oayload
	environmentDTO := r.Context().Value(validatedEnvironment{}).(dto.AddEnvironmentRequest)
	//environment := mappers.MapAddEnvironmentRequest2Environment(environmentDTO)
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	envID, err := productService.AddEnvironment(environmentDTO.ProductID, environmentDTO.Name)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.CreateSimpleEnvironmentResponse(envID, environmentDTO.Name))
	}
}

// MiddlewareValidateNewEnvironment Checks the integrity of new environment in the request and calls next if ok
func (ctx *APIContext) MiddlewareValidateNewEnvironment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		env, err := middleware.ExtractAddEnvironmentPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the user
		errs := ctx.validation.Validate(env)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the user")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedEnvironment{}, *env)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
