package cockroachdb

import (
	"context"
	"github.com/FeatureOn/api/server/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type UserRepository struct {
	cp *pgxpool.Pool
}

func newUserRepository(pool *pgxpool.Pool) UserRepository {
	return UserRepository{
		cp: pool,
	}
}

func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ur.cp.QueryRow(ctx, "select ID, Name, UserName from featureon.users where ID=$1", ID).Scan(&user.ID, &user.Name, &user.UserName); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur UserRepository) CheckUser(username string, password string) (domain.User, error) {
	user := domain.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ur.cp.QueryRow(ctx, "select ID, Name, UserName from featureon.users where UserName=$1 and Password=$2", username, password).Scan(&user.ID, &user.Name, &user.UserName); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur UserRepository) AddUser(u domain.User) error {
	id := uuid.New()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := ur.cp.Exec(ctx, "insert into featureon.users (ID, Name, UserName, Password) values ($1, $2, $3, $4)", id, u.Name, u.UserName, u.Password); err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) UpdateUser(u domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	if _, err := ur.cp.Exec(ctx, "update featureon.users set Password=$1 where ID=$2", u.Password, u.ID); err != nil {
		return err
	}
	return nil
}

func (ur UserRepository) DeleteUser(u domain.User) error {
	//TODO implement me
	panic("implement me")
}
