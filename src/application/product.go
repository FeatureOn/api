package application

import (
	"errors"

	"github.com/FeatureOn/api/domain"
	"github.com/rs/zerolog/log"
)

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
