package mappers

import (
	"errors"

	"github.com/FeatureOn/api/server/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/server/domain"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapEnvironmentFlag2EnvironmentFlagDAO maps domain EnvironmentFlag to dao EnvironmentFlagDAO
func MapEnvironmentFlag2EnvironmentFlagDAO(envFlag domain.EnvironmentFlag) (dao.EnvironmentFlagDAO, error) {
	envFlagDAO := dao.EnvironmentFlagDAO{}
	id, err := primitive.ObjectIDFromHex(envFlag.EnvironmentID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse environmentID: %s into ObjectID", envFlag.EnvironmentID)
		return envFlagDAO, errors.New("EnvironmentID format is not as expected")
	}
	envFlagDAO.EnvironmentID = id
	envFlagDAO.Flags = make([]dao.FlagDAO, 0)
	for _, flag := range envFlag.Flags {
		envFlagDAO.Flags = append(envFlagDAO.Flags, dao.FlagDAO{
			FeatureKey: flag.FeatureKey,
			Value:      flag.Value,
		})
	}
	return envFlagDAO, nil
}

// MapEnvironmentFlagDAO2EnvironmentFlag maps dao EnvironmentFlagDAO to domain EnvironmentFlag
func MapEnvironmentFlagDAO2EnvironmentFlag(envFlagDAO dao.EnvironmentFlagDAO) domain.EnvironmentFlag {
	envFlag := domain.EnvironmentFlag{
		EnvironmentID: envFlagDAO.EnvironmentID.Hex(),
	}
	for _, flagDAO := range envFlagDAO.Flags {
		envFlag.Flags = append(envFlag.Flags, mapFlagDAO2Flag(flagDAO))
	}
	return envFlag
}

func mapFlagDAO2Flag(flagDAO dao.FlagDAO) domain.Flag {
	return domain.Flag{
		FeatureKey: flagDAO.FeatureKey,
		Value:      flagDAO.Value,
	}
}
