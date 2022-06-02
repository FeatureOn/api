package cockroachdb

import "github.com/jackc/pgx/v4/pgxpool"

type HealthRepository struct {
	cp *pgxpool.Pool
}

func newHealthRepository(pool *pgxpool.Pool) HealthRepository {
	return HealthRepository{
		cp: pool,
	}
}

func (h HealthRepository) Ready() bool {
	//TODO implement me
	panic("implement me")
}
