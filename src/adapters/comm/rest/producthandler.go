package rest

import (
	"net/http"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
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
		respondWithJSON(rw, r, 200, dto.MapProduct2ProductResponse(product))
	}
}
