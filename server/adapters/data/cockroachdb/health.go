package cockroachdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// HealthRepository represent a structure that will communicate to MongoDB to accomplish health related transactions
type HealthRepository struct {
	cp     *pgxpool.Pool
	dbName string
}

func newHealthRepository(pool *pgxpool.Pool, databaseName string) HealthRepository {
	return HealthRepository{
		cp:     pool,
		dbName: databaseName,
	}
}

// Ready checks the mongodb connection
func (hr HealthRepository) Ready() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := hr.cp.Exec(ctx, fmt.Sprintf("select 1 from %s.users", hr.dbName)); err != nil {
		return false
	}
	return true
}
