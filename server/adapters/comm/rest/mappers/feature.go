package mappers

import (
	"github.com/FeatureOn/api/server/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/server/domain"
)

// MapAddFeatureRequest2Feature maps dto AddFeatureRequest to domain Feature
func MapAddFeatureRequest2Feature(feat dto.AddFeatureRequest) domain.Feature {
	return domain.Feature{
		Key:          feat.Key,
		Name:         feat.Name,
		Description:  feat.Description,
		DefaultState: feat.DefaultState,
		Active:       true,
	}
}

// MapFeature2FeatureResponse maps dto AddFeatureRequest to domain Feature
func MapFeature2FeatureResponse(feat domain.Feature) dto.FeatureResponse {
	return dto.FeatureResponse{
		Key:          feat.Key,
		Name:         feat.Name,
		Description:  feat.Description,
		DefaultState: feat.DefaultState,
		Active:       feat.Active,
	}
}
