package mappers

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/comm/rest/dto"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
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
