package rest

import (
	"context"
	"net/http"

	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/adapters/comm/rest/mappers"
	"github.com/FeatureOn/api/server/adapters/comm/rest/middleware"
	"github.com/FeatureOn/api/server/application"
	"github.com/rs/zerolog/log"
)

type validatedFeature struct{}

// swagger:route POST /feature/ Feature AddFeature
// Adds a new user to the system
// responses:
//	200: OK
//	404: errorResponse

// AddFeature creates a new environment on the system
func (apiContext *APIContext) AddFeature(rw http.ResponseWriter, r *http.Request) {
	// Get environment data from oayload
	featureDTO := r.Context().Value(validatedFeature{}).(dto.AddFeatureRequest)
	feature := mappers.MapAddFeatureRequest2Feature(featureDTO)
	productService := application.NewProductService(apiContext.productRepo, apiContext.flagRepo)
	err := productService.AddFeature(featureDTO.ProductID, feature)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapFeature2FeatureResponse(feature))
	} else {
		respondWithError(rw, r, 500, err.Error())
	}
}

// UpdateFeature creates a new environment on the system
func (apiContext *APIContext) UpdateFeature(rw http.ResponseWriter, r *http.Request) {
	// Get environment data from oayload
	featureDTO := r.Context().Value(validatedFeature{}).(dto.AddFeatureRequest)
	feature := mappers.MapAddFeatureRequest2Feature(featureDTO)
	productService := application.NewProductService(apiContext.productRepo, apiContext.flagRepo)
	err := productService.UpdateFeature(featureDTO.ProductID, feature)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapFeature2FeatureResponse(feature))
	} else {
		respondWithError(rw, r, 500, err.Error())
	}
}

// MiddlewareValidateNewFeature Checks the integrity of new feature in the request and calls next if ok
func (apiContext *APIContext) MiddlewareValidateNewFeature(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		feat, err := middleware.ExtractAddFeaturePayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the feature
		errs := apiContext.validation.Validate(feat)
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
