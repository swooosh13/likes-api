package composites

import (
	"context"
	"database/sql"
	"proj1/pkg/postgresdb"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDBComposite struct {
	Client *pgxpool.Pool
}

func NewPostgresDBComposite(ctx context.Context, host, port, username, password, database string, timeout, maxConns int) (*PostgresDBComposite, error) {
	poolConfig, err := postgresdb.NewPoolConfig(host, port, username, password, database, timeout)
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = int32(maxConns)
	c, err := postgresdb.NewConnection(poolConfig)
	if err != nil {
		return nil, err
	}

	mdb, _ := sql.Open("postgres", poolConfig.ConnString())
	err = mdb.Ping()
	if err != nil {
		return nil, err
	}

	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		return nil, err
	}

	return &PostgresDBComposite{Client: c}, nil
}
