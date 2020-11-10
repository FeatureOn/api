package domain

import (
	"crypto/sha1"
	"fmt"
	"unicode"
)

// User type defines a user of the domain
type User struct {
	ID       string
	Name     string
	UserName string
	Password string
}

// UserRepository is the interface to interact with User domain object
type UserRepository interface {
	GetUser(ID string) (User, error)
	CheckUser(username string, password string) (bool, error)
	AddUser(u User) error
	UpdateUser(u User) error
	DeleteUser(u User) error
}

// HashPassword hashes the password string in order to getting ready to store or check if it matches the stored value
func HashPassword(password string) string {
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
