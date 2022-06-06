package mappers

import (
	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
)

// CreateSimpleEnvironmentResponse creates a SimpleEnvironmentResponse with the given ID and Name
func CreateSimpleEnvironmentResponse(id string, name string) dto.SimpleEnvironmentResponse {
	return dto.SimpleEnvironmentResponse{
		ID:   id,
		Name: name,
	}
}
