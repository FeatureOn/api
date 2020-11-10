package memory

import (
	"crypto/rand"
	"fmt"
	"log"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// User object represents the user model to hold in memory
type User struct {
	ID       string
	Name     string
	UserName string
	Password string
}

var users []User

// AddUser adds a new user to the array in the memory
func AddUser(u domain.User) error {
	var user User
	user.ID = generateUUID()
	user.Name = u.Name
	user.UserName = u.UserName
	user.Password = domain.HashPassword(u.Password)
	users = append(users, user)
	return nil
}

// CheckUser checks the username & password if if matches any user frim the array
func CheckUser(username string, password string) (bool, error) {
	return false, fmt.Errorf("Not impelemnted")
}

// UpdateUser updates an existing user on the user array
func UpdateUser(u User) error {
	return fmt.Errorf("Not impelemnted")
}

// DeleteUser deletes a user from the user array
func DeleteUser(u User) error {
	return fmt.Errorf("Not impelemnted")
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
