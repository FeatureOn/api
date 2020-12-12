package application

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
)

// ProductRepository is the interface to interact with Product domain object
type ProductRepository interface {
	CheckProductNameAvailability(productID string, productName string) (bool, error)
	AddProduct(productName string) error
	UpdateProduct(productID string, productName string) error
	GetProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
	CheckEnvironmentName(productID string, envirionmentName string) (bool, error)
	AddEnvironment(productID string, environmentName string) error
	AddFeature(productID string, feat domain.Feature)
	UpdateFeature(productID string, feat domain.Feature)
	UpdateFeatureValue(productID string, environmentID string, featureID string, value bool)
	DisableFeature(productID string, feat domain.Feature)
	GetValues(productID string, environmentID string)
}

//ProductService is the struct to let outer layers to interact to the Product Applicatopn
type ProductService struct {
	productRepository ProductRepository
}

// NewProductService creates a new ProductService instance and sets its repository
func NewProductService(pr ProductRepository) ProductService {
	if pr == nil {
		panic("missing productRepository")
	}
	return ProductService{
		productRepository: pr,
	}
}

// AddProduct first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new product to the repository injected into ProductService
func (ps ProductService) AddProduct(productName string) error {
	available, err := ps.productRepository.CheckProductNameAvailability("", productName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking product name uniqueness")
		return err
	}
	if available != true {
		return errors.New("The product name is not available")
	}
	return ps.productRepository.AddProduct(productName)
}

// GetProducts returns a the list of products from the repository injected into ProductService
func (ps ProductService) GetProducts() ([]domain.Product, error) {
	return ps.productRepository.GetProducts()
}

// GetProduct returns a single product if found from the repository injected into ProductService
func (ps ProductService) GetProduct(id string) (domain.Product, error) {
	return ps.productRepository.GetProduct(id)
}

// UpdateProduct first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new product to the repository injected into ProductService
func (ps ProductService) UpdateProduct(productID string, productName string) error {
	available, err := ps.productRepository.CheckProductNameAvailability(productID, productName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking product name uniqueness")
		return err
	}
	if available != true {
		return errors.New("The product name is not available")
	}
	return ps.productRepository.UpdateProduct(productID, productName)
}
