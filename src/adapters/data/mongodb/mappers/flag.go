package mappers

import (
	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/domain"
)

// MapEnvironmentFlag2EnvironmentFlagDAO maps domain EnvironmentFlag to dao EnvironmentFlagDAO
func MapEnvironmentFlag2EnvironmentFlagDAO(envFlag domain.EnvironmentFlag) dao.EnvironmentFlag {
	envFlagDAO := dao.EnvironmentFlag{
		EnvironmentID: envFlag.EnvironmentID,
	}
	for _, flag := range envFlag.Flags {
		envFlagDAO.Flags = append(envFlagDAO.Flags, dao.Flag{
			FeatureKey: flag.FeatureKey,
			Value:      flag.Value,
		})
	}
	return envFlagDAO
}
