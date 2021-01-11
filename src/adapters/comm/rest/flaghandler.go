package rest

import (
	"net/http"

	"github.com/FeatureOn/api/adapters/comm/rest/mappers"
	"github.com/FeatureOn/api/application"
	"github.com/gorilla/mux"
)

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
