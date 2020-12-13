package application

import "dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"

// ProductRepository is the interface to interact with Product domain object
type ProductRepository interface {
	GetProductByName(productName string) (string, error)
	AddProduct(productName string) (string, error)
	UpdateProduct(productID string, productName string) error
	GetProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	GetEnvironmentByName(productID string, envirionmentName string) (string, error)
	AddEnvironment(productID string, environmentName string) (string, error)
	Updatenvironment(productID string, environmentID string, environmentName string) error
	GetEnvironments(productID string) ([]domain.Environment, error)
	GetEnvironment(productID string, environmentID string) (domain.Environment, error)
	GetFeatureByName(productID string, featureName string) (string, error)
	GetFeatureByKey(productID string, featureKey string) (string, error)
	GetFeatures(productID string) ([]domain.Feature, error)
	AddFeature(productID string, feat domain.Feature) (string, error)
	UpdateFeature(productID string, feat domain.Feature) error
	DisableFeature(productID string, feat domain.Feature) error
	UpdateFeatureValue(productID string, environmentID string, featureID string, value bool) error
}

// FlagRepository is the interface to interact with Flag domain object
type FlagRepository interface {
	AddFlag(environmentID string, FeatureID string, value bool) error
	GetFlags(environmentID string) ([]domain.Flag, error)
}

//ProductService is the struct to let outer layers to interact to the Product Applicatopn
type ProductService struct {
	productRepository ProductRepository
	flagRepository    FlagRepository
}

// NewProductService creates a new ProductService instance and sets its repository
func NewProductService(pr ProductRepository, fr FlagRepository) ProductService {
	if pr == nil {
		panic("missing productRepository")
	}
	if fr == nil {
		panic("missing flagRepository")
	}
	return ProductService{
		productRepository: pr,
		flagRepository:    fr,
	}
}
