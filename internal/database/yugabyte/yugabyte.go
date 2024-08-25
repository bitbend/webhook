package yugabyte

import (
	"context"
	"fmt"
	"github.com/driftbase/webhook/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func NewYugabyte(yugabyteConfig config.YugabyteConfig) *pgxpool.Pool {
	yugabyteUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?%s",
		yugabyteConfig.User.Username,
		yugabyteConfig.User.Password,
		yugabyteConfig.Host,
		yugabyteConfig.Port,
		yugabyteConfig.Database,
		yugabyteConfig.Options,
	)

	poolConfig, err := pgxpool.ParseConfig(yugabyteUrl)
	if err != nil {
		zap.L().Panic("error parsing yugabyte url config", zap.Error(err))
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		zap.L().Panic("error connecting to yugabyte", zap.Error(err))
	}

	return pool
}
