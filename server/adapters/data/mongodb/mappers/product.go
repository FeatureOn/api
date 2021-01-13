package mappers

import (
	"errors"

	"github.com/FeatureOn/api/server/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/server/domain"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapProductDAO2Product maps da ProductDAO to domain Product
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

// MapProduct2ProductDAO maps domain Product to dao ProductDAO
func MapProduct2ProductDAO(product domain.Product) (dao.ProductDAO, error) {
	productDAO := dao.ProductDAO{}
	id, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ProductID %s into bson ObjectID", product.ID)
		return productDAO, errors.New("ProductID format is not as expected")
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
	return productDAO, nil
}

// MapFeature2FeatureDAO maps domain Feature to dao FeatureDAO
func MapFeature2FeatureDAO(feat domain.Feature) dao.FeatureDAO {
	return dao.FeatureDAO{
		Key:          feat.Key,
		Name:         feat.Name,
		Description:  feat.Description,
		DefaultState: feat.DefaultState,
		Active:       feat.Active,
	}
}

// MapFeatureDAO2Feature maps dao FeatureDAO to domain Feature
func MapFeatureDAO2Feature(feat dao.FeatureDAO) domain.Feature {
	return domain.Feature{
		Key:          feat.Key,
		Name:         feat.Name,
		Description:  feat.Description,
		DefaultState: feat.DefaultState,
		Active:       feat.Active,
	}
}
