package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/synt4xer/govest/config"
)

func ProvideDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	uri := cfg.Database.URI

	db, err := sqlx.ConnectContext(ctx, "postgres", uri)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(cfg.Database.MaxIdleConnsLifetime)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
