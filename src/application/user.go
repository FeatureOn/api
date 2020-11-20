package application

import (
	"crypto/sha1"
	"fmt"
	"unicode"

	"dev.azure.com/serdarkalayci-github/Toggler/_git/toggler-api/domain"
)

// UserRepository is the interface to interact with User domain object
type UserRepository interface {
	GetUser(ID string) (domain.User, error)
	CheckUser(username string, password string) (domain.User, error)
	AddUser(u domain.User) error
	UpdateUser(u domain.User) error
	DeleteUser(u domain.User) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(ur UserRepository) UserService {
	if ur == nil {
		panic("missing userRepository")
	}
	return UserService{
		userRepository: ur,
	}
}

func (us UserService) GetUser(ID string) (domain.User, error) {
	return us.userRepository.GetUser(ID)
}

func (us UserService) CheckUser(username string, password string) (domain.User, error) {
	return us.userRepository.CheckUser(username, hashPassword(password))
}

func (us UserService) AddUser(u domain.User) error {
	u.Password = hashPassword(u.Password)
	return us.userRepository.AddUser(u)
}

func (us UserService) UpdateUser(u domain.User) error {
	return us.userRepository.UpdateUser(u)
}

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
