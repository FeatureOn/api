package mappers

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// MapProduct2ProductResponse maps domain Product to dto ProductResponse
func MapProduct2ProductResponse(products []domain.Product) []dto.ProductResponse {
	productResponses := make([]dto.ProductResponse, 0)
	for _, product := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:   product.ID,
			Name: product.Name,
		})
	}
	return productResponses
}

// MapProduct2ProductDetailResponse maps domain Product to ProductDetailResponse
func MapProduct2ProductDetailResponse(product domain.Product) dto.ProductDetailResponse {
	productDetailResponse := dto.ProductDetailResponse{
		ID:   product.ID,
		Name: product.Name,
	}
	for _, env := range product.Environments {
		productDetailResponse.Environments = append(productDetailResponse.Environments, dto.EnvironmentResponse{
			ID:   env.ID,
			Name: env.Name,
		})
	}
	for _, feat := range product.Features {
		productDetailResponse.Features = append(productDetailResponse.Features, dto.FeatureResponse{
			Name:         feat.Name,
			Key:          feat.Key,
			Description:  feat.Description,
			DefaultState: feat.DefaultState,
			Active:       feat.Active,
		})
	}
	return productDetailResponse
}
