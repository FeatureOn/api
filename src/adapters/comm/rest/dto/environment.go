package dto

// AddEnvironmentRequest type defines a model for adding an environment
type AddEnvironmentRequest struct {
	ProductID string `json:"productID" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

// UpdateEnvironmentRequest type defines a model for updating an environment
type UpdateEnvironmentRequest struct {
	ProductID     string `json:"productID" validate:"required"`
	EnvironmentID string `json:"environmentID" validate:"required"`
	Name          string `json:"name" validate:"required"`
}

// SimpleEnvironmentResponse type defines a model for returning simple environment data
type SimpleEnvironmentResponse struct {
	ID   string `json:"environmentID"`
	Name string `json:"name"`
}
