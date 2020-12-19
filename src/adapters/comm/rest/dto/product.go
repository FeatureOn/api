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

// MapProduct2ProductDetailResponse maps domain Product to ProductDetailResponse
func MapProduct2ProductDetailResponse(product domain.Product) ProductDetailResponse {
	productDetailResponse := ProductDetailResponse{
		ID:   product.ID,
		Name: product.Name,
	}
	for _, env := range product.Environments {
		productDetailResponse.Environments = append(productDetailResponse.Environments, EnvironmentResponse{
			ID:   env.ID,
			Name: env.Name,
		})
	}
	for _, feat := range product.Features {
		productDetailResponse.Features = append(productDetailResponse.Features, FeatureResponse{
			Name:         feat.Name,
			Key:          feat.Key,
			Description:  feat.Description,
			DefaultState: feat.DefaultState,
			Active:       feat.Active,
		})
	}
	return productDetailResponse
}

// ProductDetailResponse is the full body response of a product
type ProductDetailResponse struct {
	ID           string
	Name         string
	Features     []FeatureResponse
	Environments []EnvironmentResponse
}

// FeatureResponse is a basic flag (as for now) holding a key within a project and its default state
type FeatureResponse struct {
	Name         string
	Key          string
	Description  string
	DefaultState bool
	Active       bool
}

// EnvironmentResponse is a struct that will hold the collection of flags for each of product's deployment
type EnvironmentResponse struct {
	ID   string
	Name string
}
