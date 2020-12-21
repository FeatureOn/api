package mappers

import (
	"github.com/FeatureOn/api/adapters/comm/rest/dto"
	"github.com/FeatureOn/api/domain"
)

// MapUserRequest2User maps dto UserRequest to domain User
func MapUserRequest2User(ur dto.UserRequest) domain.User {
	return domain.User{
		ID:       ur.ID,
		Name:     ur.Name,
		UserName: ur.UserName,
		Password: ur.Password,
	}
}

// MapUser2UserResponse maps domain User to dto UserResponse
func MapUser2UserResponse(u domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.UserName,
	}
}
