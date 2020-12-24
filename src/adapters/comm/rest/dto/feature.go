package dto

// AddFeatureRequest type defines a model for adding a feature
type AddFeatureRequest struct {
	ProductID    string `json:"productID" validare:"required"`
	Name         string `json:"name" validate:"required"`
	Key          string `json:"key" validate:"required"`
	Description  string `json:"description"`
	DefaultState bool   `json:"defaultstate"`
}

// UpdateFeatureRequest type defines a model for updating a feature
type UpdateFeatureRequest struct {
	ProductID    string `json:"productID" validare:"required"`
	Name         string `json:"name"`
	Key          string `json:"key" validate:"required"`
	Description  string `json:"description"`
	DefaultState bool   `json:"defaultstate"`
	Active       bool   `json:"active"`
}
