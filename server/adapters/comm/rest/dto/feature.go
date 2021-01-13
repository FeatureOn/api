package dto

// AddFeatureRequest type defines a model for adding a feature
type AddFeatureRequest struct {
	ProductID    string `json:"productID" validare:"required"`
	Name         string `json:"name" validate:"required"`
	Key          string `json:"key" validate:"required"`
	Description  string `json:"description" validate:"required"`
	DefaultState bool   `json:"defaultstate"` // ToDo: Validator only accepts true value here if required tag is entered.
}
