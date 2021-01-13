package mappers

import (
	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/domain"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapUser2UserDAO maps domain User to DAO UserDAO
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

// MapUser2NewUserDAO maps domain User to dao User
func MapUser2NewUserDAO(u domain.User) dao.UserDAO {
	userDAO := dao.UserDAO{}
	userDAO.ID = primitive.NewObjectID()
	userDAO.Name = u.Name
	userDAO.UserName = u.UserName
	userDAO.Password = u.Password
	return userDAO
}

// MapUserDAO2User maps dao User to domain User
func MapUserDAO2User(u dao.UserDAO) domain.User {
	user := domain.User{}
	user.ID = u.ID.Hex()
	user.Name = u.Name
	user.UserName = u.UserName
	user.Password = u.Password
	return user
}
