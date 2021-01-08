package mappers

import (
	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/domain"
)

// MapEnvironmentFlag2EnvironmentFlagDAO maps domain EnvironmentFlag to dao EnvironmentFlagDAO
func MapEnvironmentFlag2EnvironmentFlagDAO(envFlag domain.EnvironmentFlag) dao.EnvironmentFlag {
	envFlagDAO := dao.EnvironmentFlag{}
	envFlagDAO.EnvironmentID = envFlag.EnvironmentID
	envFlagDAO.Flags = make([]dao.FlagDAO, 0)
	for _, flag := range envFlag.Flags {
		envFlagDAO.Flags = append(envFlagDAO.Flags, dao.FlagDAO{
			FeatureKey: flag.FeatureKey,
			Value:      flag.Value,
		})
	}
	return envFlagDAO
}
