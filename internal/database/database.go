package database

import (
	"context"
	"fmt"
	"github.com/bitbend/webhook/internal/config"
	"github.com/bitbend/webhook/internal/database/postgres"
	"github.com/bitbend/webhook/internal/database/yugabyte"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(databaseConfig config.DatabaseConfig) *Database {
	var pool *pgxpool.Pool

	switch databaseConfig.Provider {
	case "postgres":
		pool = postgres.NewPostgres(databaseConfig.Postgres)
	case "yugabyte":
		pool = yugabyte.NewYugabyte(databaseConfig.Yugabyte)
	default:
		zap.L().Panic("unsupported database provider", zap.String("provider", databaseConfig.Provider))
	}

	return &Database{
		pool: pool,
	}
}

func (db *Database) WithTxn(ctx context.Context, txOptions pgx.TxOptions, fn func(tx pgx.Tx, ctx context.Context) error) error {
	tx, err := db.pool.BeginTx(ctx, txOptions)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				err = fmt.Errorf("failed to rollback transaction: %v, original error: %w", rbErr, err)
			}
		}
	}()

	err = fn(tx, ctx)

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
