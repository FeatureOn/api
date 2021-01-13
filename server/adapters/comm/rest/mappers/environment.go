package mappers

import (
	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/domain"
)

// MapAddEnvironmentRequest2Environment maps dto AddEnvironmentRequest to domain Environment
func MapAddEnvironmentRequest2Environment(environment dto.AddEnvironmentRequest) domain.Environment {
	return domain.Environment{
		Name: environment.Name,
	}
}

// CreateSimpleEnvironmentResponse creates a SimpleEnvironmentResponse with the given ID and Name
func CreateSimpleEnvironmentResponse(id string, name string) dto.SimpleEnvironmentResponse {
	return dto.SimpleEnvironmentResponse{
		ID:   id,
		Name: name,
	}
}
