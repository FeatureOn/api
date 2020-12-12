package application

import (
	"errors"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
)

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
	AddFeature(productID string, feat domain.Feature) (string, error)
	UpdateFeature(productID string, feat domain.Feature) error
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
func (ps ProductService) AddProduct(productName string) (string, error) {
	existingID, err := ps.productRepository.GetProductByName(productName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking product name uniqueness")
		return "", err
	}
	if existingID != "" {
		return "", errors.New("The product name is not available")
	}
	return ps.productRepository.AddProduct(productName)
}

// UpdateProduct first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new product to the repository injected into ProductService
func (ps ProductService) UpdateProduct(productID string, productName string) error {
	existingID, err := ps.productRepository.GetProductByName(productName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking product name uniqueness")
		return err
	}
	if existingID != productID {
		return errors.New("The product name is not available")
	}
	return ps.productRepository.UpdateProduct(productID, productName)
}

// GetProducts returns a the list of products from the repository injected into ProductService
func (ps ProductService) GetProducts() ([]domain.Product, error) {
	return ps.productRepository.GetProducts()
}

// GetProduct returns a single product if found from the repository injected into ProductService
func (ps ProductService) GetProduct(id string) (domain.Product, error) {
	return ps.productRepository.GetProduct(id)
}

// AddEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it adds a new environment on the product to the repository injected into ProductService
func (ps ProductService) AddEnvironment(productID string, environmentName string) (string, error) {
	existingID, err := ps.productRepository.GetEnvironmentByName(productID, environmentName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking environment name uniqueness")
		return "", err
	}
	if existingID != productID {
		return "", errors.New("The environment name is not available")
	}
	return ps.productRepository.AddEnvironment(productID, environmentName)
}

// UpdateEnvironment first checks the availability of the name because the system should not allow the same name used twice
// if the name is unique, it updates the existing environment on the product to the repository injected into ProductService
func (ps ProductService) UpdateEnvironment(productID string, environmentID string, environmentName string) error {
	existingID, err := ps.productRepository.GetEnvironmentByName(productID, environmentName)
	if err != nil {
		log.Error().Err(err).Msg("Error checking environment name uniqueness")
		return err
	}
	if existingID != environmentName {
		return errors.New("The environment name is not available")
	}
	return ps.productRepository.Updatenvironment(productID, environmentID, environmentName)
}

// GetEnvironments returns a the list of environments defined for a product from the repository injected into ProductService
func (ps ProductService) GetEnvironments(productID string) ([]domain.Environment, error) {
	return ps.productRepository.GetEnvironments(productID)
}

// GetEnvironment returns an environment defined for a product from the repository injected into ProductService
func (ps ProductService) GetEnvironment(productID string, environmentID string) (domain.Environment, error) {
	return ps.productRepository.GetEnvironment(productID, environmentID)
}

// AddFeature first checks the uniqueness of Feature@s Name and Key because the system should not allow Name and Key used twice
func (ps ProductService) AddFeature(productID string, feat domain.Feature) (string, error) {
	existingID, err := ps.productRepository.GetFeatureByName(productID, feat.Name)
	if err != nil {
		log.Error().Err(err).Msg("Error checking feature name uniqueness")
		return "", err
	}
	if existingID != "" {
		return "", errors.New("The freature name is not available")
	}
	existingID, err = ps.productRepository.GetFeatureByKey(productID, feat.Key)
	if err != nil {
		log.Error().Err(err).Msg("Error checking feature key uniqueness")
		return "", err
	}
	if existingID != "" {
		return "", errors.New("The freature key is not available")
	}
	return ps.productRepository.AddFeature(productID, feat)
}
