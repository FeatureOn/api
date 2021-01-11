package mappers

import (
	"github.com/FeatureOn/api/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/domain"
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