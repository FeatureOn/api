package application

import "github.com/FeatureOn/api/domain"

// ProductRepository is the interface to interact with Product domain object
type ProductRepository interface {
	GetProductByName(productName string) (string, error)
	AddProduct(productName string) (string, error)
	UpdateProduct(product string, productName string) error
	GetProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	AddEnvironment(product domain.Product, environmentName string, environmentFlag domain.EnvironmentFlag) (string, error)
	UpdateEnvironment(product domain.Product, environmentID string, environmentName string) error
	AddFeature(product domain.Product, feature domain.Feature, envFlags []domain.EnvironmentFlag) error
	UpdateFeature(product domain.Product, feature domain.Feature) error
	ToggleFeatureState(product domain.Product, featureID string, newState bool) error
}

// FlagRepository is the interface to interact with Flag domain object
type FlagRepository interface {
	GetFlags(environmentID string) ([]domain.Flag, error)
	UpdateFlag(productID string, environmentID string, featureID string, value bool) error
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
