package application

import (
	"crypto/sha1"
	"fmt"
	"unicode"

	"github.com/FeatureOn/api/server/domain"
)

// UserRepository is the interface to interact with User domain object
type UserRepository interface {
	GetUser(ID string) (domain.User, error)
	CheckUser(username string, password string) (domain.User, error)
	AddUser(u domain.User) error
	UpdateUser(u domain.User) error
	DeleteUser(u domain.User) error
}

//UserService is the struct to let outer layers to interact to the User Applicatopn
type UserService struct {
	userRepository UserRepository
}

// NewUserService creates a new UserService instance and sets its repository
func NewUserService(ur UserRepository) UserService {
	if ur == nil {
		panic("missing userRepository")
	}
	return UserService{
		userRepository: ur,
	}
}

// GetUser simply returns a single user or an error that is returned from the repository
func (us UserService) GetUser(ID string) (domain.User, error) {
	return us.userRepository.GetUser(ID)
}

// CheckUser checks if the username and password maches any from the repository by first hashing its password, returns error if none found
func (us UserService) CheckUser(username string, password string) (domain.User, error) {
	return us.userRepository.CheckUser(username, hashPassword(password))
}

// AddUser adds a new user to the repository by first hashing its password
func (us UserService) AddUser(u domain.User) error {
	u.Password = hashPassword(u.Password)
	return us.userRepository.AddUser(u)
}

// UpdateUser updates a single user on the repository, returns error if repository returns one
func (us UserService) UpdateUser(u domain.User) error {
	return us.userRepository.UpdateUser(u)
}

// DeleteUser deletes a single user from the repository, returns error if repository returns one
func (us UserService) DeleteUser(u domain.User) error {
	return us.userRepository.DeleteUser(u)
}

// HashPassword hashes the password string in order to getting ready to store or check if it matches the stored value
func hashPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return string(h.Sum(nil))
}

func checkPassword(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"upper case": {unicode.Upper, unicode.Title},
		"lower case": {unicode.Lower},
		"numeric":    {unicode.Number, unicode.Digit},
		"special":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return fmt.Errorf("password must have at least one %s character", name)
	}
	return nil
}
