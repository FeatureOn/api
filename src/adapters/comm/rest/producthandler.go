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

type validatedProduct struct{}

// swagger:route GET /product/{id} Product GetProduct
// Return the product if found
// responses:
//	200: OK
//	404: errorResponse

// GetProduct gets a single product if found
func (ctx *APIContext) GetProduct(rw http.ResponseWriter, r *http.Request) {
	// parse the Rating id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id := vars["id"]
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	product, err := productService.GetProduct(id)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapProduct2ProductDetailResponse(product))
	}
}

// swagger:route GET /product Products GetProducts
// Return the list of products if found
// responses:
//	200: OK
//	404: errorResponse

// GetProducts gets a list of products if found
func (ctx *APIContext) GetProducts(rw http.ResponseWriter, r *http.Request) {
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	products, err := productService.GetProducts()
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.MapProducts2ProductResponses(products))
	}
}

// swagger:route POST /product AddProduct AddProduct
// Adds a new product to the system
// responses:
//	200: OK
//	404: errorResponse

// AddProduct adds a new product to the system
func (ctx *APIContext) AddProduct(rw http.ResponseWriter, r *http.Request) {
	// Get product data from payload
	productDTO := r.Context().Value(validatedProduct{}).(dto.AddProductRequest)
	//environment := mappers.MapAddEnvironmentRequest2Environment(environmentDTO)
	productService := application.NewProductService(ctx.productRepo, ctx.flagRepo)
	prodID, err := productService.AddProduct(productDTO.Name)
	if err == nil {
		respondWithJSON(rw, r, 200, mappers.CreateSimpleProductResponse(prodID, productDTO.Name))
	}
}

// MiddlewareValidateNewProduct Checks the integrity of new product in the request and calls next if ok
func (ctx *APIContext) MiddlewareValidateNewProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		env, err := middleware.ExtractAddProductPayload(r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// validate the product
		errs := ctx.validation.Validate(env)
		if errs != nil && len(errs) != 0 {
			log.Error().Err(errs[0]).Msg("Error validating the product")

			// return the validation messages as an array
			respondWithJSON(rw, r, http.StatusUnprocessableEntity, errs.Errors())
			return
		}

		// add the rating to the context
		ctx := context.WithValue(r.Context(), validatedProduct{}, *env)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
