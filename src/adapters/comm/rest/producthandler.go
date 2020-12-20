package rest

import (
	"net/http"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/mappers"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"
	"github.com/gorilla/mux"
)

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
		respondWithJSON(rw, r, 200, mappers.MapProduct2ProductResponse(products))
	}
}
