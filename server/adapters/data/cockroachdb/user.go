package cockroachdb

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/FeatureOn/api/server/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// UserRepository represent a structure that will communicate to MongoDB to accomplish user related transactions
type UserRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newUserRepository(pool *pgxpool.Pool, databaseName string) UserRepository {
	return UserRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

// GetUser returns one user with the given ID if it exists in the array, returns not found error otherwise
func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ur.cp.QueryRow(ctx, fmt.Sprintf("select ID, Name, UserName from %s.users where ID=$1", ur.dbName), ID).Scan(&user.ID, &user.Name, &user.UserName); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// CheckUser checks the username & password if it matches any user from the array
func (ur UserRepository) CheckUser(username string, password string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ur.cp.QueryRow(ctx, fmt.Sprintf("select ID, Name, UserName from %s.users where UserName=$1 and Password=$2", ur.dbName), username, password).Scan(&user.ID, &user.Name, &user.UserName); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// AddUser adds a new user to the array in the memory
func (ur UserRepository) AddUser(u domain.User) error {
	id := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	b := []byte(u.Password)
	encodedString := hex.EncodeToString(b)
	fmt.Println(encodedString)
	if _, err := ur.cp.Exec(ctx, fmt.Sprintf("insert into %s.users (ID, Name, UserName, Password) values ($1, $2, $3, $4)", ur.dbName), id, u.Name, u.UserName, u.Password); err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user on the user array
func (ur UserRepository) UpdateUser(u domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	if _, err := ur.cp.Exec(ctx, fmt.Sprintf("update %s.users set Password=$1 where ID=$2", ur.dbName), u.Password, u.ID); err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the user array
func (ur UserRepository) DeleteUser(u domain.User) error {
	//TODO implement me
	panic("implement me")
}
