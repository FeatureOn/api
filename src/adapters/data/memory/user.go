package memory

import (
	"crypto/rand"
	"fmt"
	"log"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/application"
	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// // User object represents the user model to hold in memory
// type User struct {
// 	ID       string
// 	Name     string
// 	UserName string
// 	Password string
// }

var users []domain.User

type UserRepository struct{}

func newUserRepository() UserRepository {
	createInitialUsers()
	return UserRepository{}
}

func createInitialUsers() {
	user := domain.User{}
	user.ID = "b2bf3967-1991-fbf1-0d3e-93222d2a4050"
	user.Name = "First User"
	user.UserName = "firstu"
	user.Password = application.HashPassword("firstp")
	users = append(users, user)
	user.ID = "1b644ef3-1fd3-ad9b-54ad-cbe9eff5bbfb"
	user.Name = "Second User"
	user.UserName = "secondu"
	user.Password = application.HashPassword("secondp")
	users = append(users, user)
}

// GetUser returns one user with the given ID if it exists in the array, returns not found error otherwise
func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	for _, user := range users {
		if user.ID == ID {
			return user, nil
		}
	}
	return domain.User{}, fmt.Errorf("Not found")
}

// AddUser adds a new user to the array in the memory
func (ur UserRepository) AddUser(u domain.User) error {
	var user domain.User
	user.ID = generateUUID()
	user.Name = u.Name
	user.UserName = u.UserName
	user.Password = application.HashPassword(u.Password)
	users = append(users, user)
	return nil
}

// CheckUser checks the username & password if if matches any user frim the array
func (ur UserRepository) CheckUser(username string, password string) (domain.User, error) {
	for _, user := range users {
		if user.UserName == username && user.Password == password {
			return user, nil
		}
	}
	return domain.User{}, fmt.Errorf("Not found")
}

// UpdateUser updates an existing user on the user array
func (ur UserRepository) UpdateUser(u domain.User) error {
	return fmt.Errorf("Not impelemented")
}

// DeleteUser deletes a user from the user array
func (ur UserRepository) DeleteUser(u domain.User) error {
	return fmt.Errorf("Not impelemented")
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return string(uuid)
}
