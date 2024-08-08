package psql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
	"warehouse/config"
)

type Postgres struct {
	Conn *pgx.Conn
}

func New(conf *config.Config) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, conf.PG.DSN)
	if err != nil {
		return nil, err
	}

	return &Postgres{Conn: conn}, nil
}
