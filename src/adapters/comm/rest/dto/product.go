package dto

// ProductResponse represents a product type returned by a rest service
type ProductResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ProductDetailResponse is the full body response of a product
type ProductDetailResponse struct {
	ID           string
	Name         string
	Features     []FeatureResponse
	Environments []EnvironmentResponse
}

// FeatureResponse is a basic flag (as for now) holding a key within a project and its default state
type FeatureResponse struct {
	Name         string
	Key          string
	Description  string
	DefaultState bool
	Active       bool
}

// EnvironmentResponse is a struct that will hold the collection of flags for each of product's deployment
type EnvironmentResponse struct {
	ID   string
	Name string
}
