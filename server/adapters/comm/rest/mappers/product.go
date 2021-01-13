package mappers

import (
	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/domain"
)

// MapProducts2ProductResponses maps domain Product to dto ProductResponse
func MapProducts2ProductResponses(products []domain.Product) []dto.ProductResponse {
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

// MapAddProductRequest2Request maps dto AddProductRequest to domain Product
func MapAddProductRequest2Request(product dto.AddProductRequest) domain.Product {
	return domain.Product{
		Name: product.Name,
	}
}

// MapUpdateProductRequest2Request maps dto UpdateProductRequest to domain Product
func MapUpdateProductRequest2Request(product dto.UpdateProductRequest) domain.Product {
	return domain.Product{
		ID:   product.ID,
		Name: product.Name,
	}
}

// CreateSimpleProductResponse creates a ProductResponse with the given ID and Name
func CreateSimpleProductResponse(id string, name string) dto.ProductResponse {
	return dto.ProductResponse{
		ID:   id,
		Name: name,
	}
}
