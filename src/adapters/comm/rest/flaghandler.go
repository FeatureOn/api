package rest

import (
	"context"
	"net/http"

	"github.com/FeatureOn/api/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/adapters/comm/rest/mappers"
	middleware "github.com/FeatureOn/api/adapters/comm/rest/middleware"
	"github.com/FeatureOn/api/application"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type validatedFlag struct{}

// swagger:route GET /flag/ Flag GetFlags
// Gets all the flags within an environment
// responses:
//	200: OK
//	404: errorResponse

// GetFlags gets all flags within an environment
func (ctx *APIContext) GetFlags(rw http.ResponseWriter, r *http.Request) {
	// parse the environment id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	envFlag, err := productService.GetFlags(id)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapEnvironmentFlag2EnvironmentFlagResponse(envFlag))
	} else {
		respondWithError(rw, r, 500, err.Error())
	}
}

// swagger:route PUT /flag UpdateFlag UpdateFlag
// Updates an existing flag's value on the system
// responses:
//	200: OK
//	404: errorResponse

// UpdateProducUpdateFlagt adds a new product to the system
func (ctx *APIContext) UpdateFlag(rw http.ResponseWriter, r *http.Request) {
	// Get product data from payload
	updateFlagDTO := r.Context().Value(validatedFlag{}).(dto.UpdateFlagRequest)
	//environment := mappers.MapAddEnvironmentRequest2Environment(environmentDTO)
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	err := productService.UpdateFlagValue(updateFlagDTO.EnvironmentID, updateFlagDTO.FeatureKey, updateFlagDTO.Value)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.CreateFlagResponse(updateFlagDTO.FeatureKey, updateFlagDTO.Value))
	} else {
		respondWithError(rw, r, 500, err.Error())
	}
}

// MiddlewareValidateUpdateFlag validates the input of UpdateFlagRequest
func (ctx *APIContext) MiddlewareValidateUpdateFlag(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		updFlag, err := middleware.ExtractUpdateFlagPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the product
		errs := ctx.validation.Validate(updFlag)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the UpdateFlagRequest")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedFlag{}, *updFlag)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
