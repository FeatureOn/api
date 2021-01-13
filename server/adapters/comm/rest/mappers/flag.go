package mappers

import (
	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/domain"
)

// MapEnvironmentFlag2EnvironmentFlagResponse maps domain EnvironmentFlag to dto EnvironmentFlagResponse
func MapEnvironmentFlag2EnvironmentFlagResponse(envFlag domain.EnvironmentFlag) dto.EnvironmentFlagResponse {
	envFlagDTO := dto.EnvironmentFlagResponse{
		EnvironmentID: envFlag.EnvironmentID,
	}
	for _, flag := range envFlag.Flags {
		envFlagDTO.Flags = append(envFlagDTO.Flags, mapFlag2FlagResponse(flag))
	}
	return envFlagDTO
}

func mapFlag2FlagResponse(flag domain.Flag) dto.FlagResponse {
	return dto.FlagResponse{
		FeatureKey: flag.FeatureKey,
		Value:      flag.Value,
	}
}

// CreateFlagResponse creates a new FlagResponse from given parameters and returns its
func CreateFlagResponse(featureKey string, value bool) dto.FlagResponse {
	return dto.FlagResponse{
		FeatureKey: featureKey,
		Value:      value,
	}
}
