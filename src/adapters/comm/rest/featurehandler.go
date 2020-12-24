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

type validatedFeature struct{}

// swagger:route POST /feature/ Feature AddFeature
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddFeature creates a new environment on the system
func (ctx *APIContext) AddFeature(rw http.ResponseWriter, r *http.Request) {
	// Get environment data from oayload
	featureDTO := r.Context().Value(validatedFeature{}).(dto.AddFeatureRequest)
	feature := mappers.MapAddFeatureRequest2Feature(featureDTO)
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	_, err := productService.AddFeature(featureDTO.ProductID, feature)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapFeature2FeatureResponse(feature))
	}
}

// MiddlewareValidateNewFeature Checks the integrity of new feature in the request and calls next if ok
func (ctx *APIContext) MiddlewareValidateNewFeature(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		feat, err := middleware.ExtractAddFeaturePayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the feature
		errs := ctx.validation.Validate(feat)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the feature")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedFeature{}, *feat)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
