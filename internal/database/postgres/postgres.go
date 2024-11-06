package postgres

import (
	"context"
	"fmt"
	"github.com/bitbend/webhook/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewPostgres(postgresConfig config.PostgresConfig) *pgxpool.Pool {
	postgresUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		postgresConfig.User.Username,
		postgresConfig.User.Password,
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.Database,
		postgresConfig.Options,
	)

	poolConfig, err := pgxpool.ParseConfig(postgresUrl)
	if err != nil {
		zap.L().Panic("error parsing postgres url config", zap.Error(err))
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		zap.L().Panic("error connecting to postgres", zap.Error(err))
	}

	return pool
}
