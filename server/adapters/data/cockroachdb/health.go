package cockroachdb

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type HealthRepository struct {
	cp *pgxpool.Pool
}

func newHealthRepository(pool *pgxpool.Pool) HealthRepository {
	return HealthRepository{
		cp: pool,
	}
}

func (hr HealthRepository) Ready() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := hr.cp.Exec(ctx, "select 1 from featureon.users"); err != nil {
		return false
	}
	return true
}
