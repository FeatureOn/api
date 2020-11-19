package mongodb

import (
	"fmt"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository represent a structure that will communicate to MongoDB to accomplish user related transactions
type UserRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newUserRepository(client *mongo.Client, databaseName string) UserRepository {
	return UserRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetUser returns one user with the given ID if it exists in the array, returns not found error otherwise
func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Not impelemented")
}

// AddUser adds a new user to the array in the memory
func (ur UserRepository) AddUser(u domain.User) error {
	return fmt.Errorf("Not impelemented")
}

// CheckUser checks the username & password if if matches any user frim the array
func (ur UserRepository) CheckUser(username string, password string) (domain.User, error) {
	return domain.User{}, fmt.Errorf("Not impelemented")
}

// UpdateUser updates an existing user on the user array
func (ur UserRepository) UpdateUser(u domain.User) error {
	return fmt.Errorf("Not impelemented")
}

// DeleteUser deletes a user from the user array
func (ur UserRepository) DeleteUser(u domain.User) error {
	return fmt.Errorf("Not impelemented")
}
