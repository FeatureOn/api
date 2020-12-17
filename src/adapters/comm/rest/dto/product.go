package dto

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// ProductResponse represents a product type returned by a rest service
type ProductResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MapProduct2ProductResponse maps domain Product to dto ProductResponse
func MapProduct2ProductResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		ID:   product.ID,
		Name: product.Name,
	}
}
