package dto

// AddEnvironmentRequest type defines a model for adding or updating an environment
type AddEnvironmentRequest struct {
	ProductID string `json:"productID" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

// SimpleEnvironmentResponse type defines a model for returning simple environment data
type SimpleEnvironmentResponse struct {
	ID   string `json:"environmentID"`
	Name string `json:"name"`
}
