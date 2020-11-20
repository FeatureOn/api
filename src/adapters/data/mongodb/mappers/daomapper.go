package mappers

import (
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/adapters/data/mongodb/dao"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapUser2UserDAO(u domain.User) (dao.UserDAO, error) {
	userDAO := dao.UserDAO{}
	id, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse ObjectID of UserID: %s", u.ID)
	}
	userDAO.ID = id
	userDAO.Name = u.Name
	userDAO.UserName = u.UserName
	userDAO.Password = u.Password
	return userDAO, err
}

func MapUser2NewUserDAO(u domain.User) dao.UserDAO {
	userDAO := dao.UserDAO{}
	userDAO.ID = primitive.NewObjectID()
	userDAO.Name = u.Name
	userDAO.UserName = u.UserName
	userDAO.Password = u.Password
	return userDAO
}

func MapUserDAO2User(u dao.UserDAO) domain.User {
	user := domain.User{}
	user.ID = u.ID.Hex()
	user.Name = u.Name
	user.UserName = u.UserName
	user.Password = u.Password
	return user
}
