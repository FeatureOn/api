package dto

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// UserRequest type defines a model for adding or updating an user
type UserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UserResponse type defines a model for returning an user shy of its password
type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

// LoginRequest type defines a model for getting an user's data for login operation
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func MapUserRequest2User(ur UserRequest) domain.User {
	return domain.User{
		ID:       ur.ID,
		Name:     ur.Name,
		UserName: ur.UserName,
		Password: ur.Password,
	}
}

func MapUser2UserResponse(u domain.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.UserName,
	}
}
