package mappers

import (
	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/domain"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapProductDAO2Product(productDAO dao.ProductDAO) domain.Product {
	product := domain.Product{
		ID:   productDAO.ID.Hex(),
		Name: productDAO.Name,
	}
	for _, env := range productDAO.Environments {
		product.Environments = append(product.Environments, domain.Environment{
			ID:   env.ID.Hex(),
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

func MapProduct2ProductDAO(product domain.Product) dao.ProductDAO {
	productDAO := dao.ProductDAO{}
	id, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ProductID %s into bson ObjectID", product.ID)
	}
	productDAO.ID = id
	productDAO.Name = product.Name

	for _, env := range product.Environments {
		id, _ := primitive.ObjectIDFromHex(env.ID)
		productDAO.Environments = append(productDAO.Environments, dao.EnvironmentDAO{
			ID:   id,
			Name: env.Name,
		})
	}
	for _, feat := range product.Features {
		productDAO.Features = append(productDAO.Features, dao.FeatureDAO{
			Name:         feat.Name,
			Key:          feat.Key,
			Description:  feat.Description,
			DefaultState: feat.DefaultState,
			Active:       feat.Active,
		})
	}
	return productDAO
}
