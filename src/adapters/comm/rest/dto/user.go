package dto

// UserRequest type defines a model for adding or updating an user
type UserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse type defines a model for returning an user shy of its password
type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

// LoginRequest type defines a model for getting an user's data for login operation
type LoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
