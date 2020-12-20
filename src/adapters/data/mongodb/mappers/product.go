package mappers

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/mongodb/dao"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

func MapProductDAO2Product(productDAO dao.ProductDAO) domain.Product {
	product := domain.Product{
		ID:   productDAO.ID.Hex(),
		Name: productDAO.Name,
	}
	for _, env := range productDAO.Environments {
		product.Environments = append(product.Environments, domain.Environment{
			ID:   env.ID,
			Name: env.Name,
		})
	}
	for _, feat := range productDAO.Features {
		product.Features = append(product.Features, domain.Feature{
			Name:         feat.Name,
			Key:          feat.Key,
			Description:  feat.Description,
			DefaultState: feat.DefaultState,
			Active:       feat.Active,
		})
	}
	return product
}
