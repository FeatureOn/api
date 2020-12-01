package application

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// ProductRepository is the interface to interact with Product domain object
type ProductRepository interface {
	AddProduct(p domain.Product) error
	GetProducts() ([]domain.Product, error)
	GetProduct(id string) (domain.Product, error)
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

// AddProduct adds a new product to the repository injected into ProductService
func (ps ProductService) AddProduct(p domain.Product) error {
	return ps.productRepository.AddProduct(p)
}

// GetProducts returns a the list of products from the repository injected into ProductService
func (ps ProductService) GetProducts() ([]domain.Product, error) {
	return ps.productRepository.GetProducts()
}

// GetProduct returns a single product if found from the repository injected into ProductService
func (ps ProductService) GetProduct(id string) (domain.Product, error) {
	return ps.productRepository.GetProduct(id)
}
