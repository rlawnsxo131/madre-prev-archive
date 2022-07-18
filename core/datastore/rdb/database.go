package rdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rs/zerolog"
)

const (
	KEY_DATABASE_CTX = "KEY_DATABASE_CTX"
)

var (
	pool         *pgxpool.Pool
	onceDatabase sync.Once
)

func InitDatabase() (*pgxpool.Pool, error) {
	var err error

	onceDatabase.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			env.DatabaseHost(),
			env.DatabasePort(),
			env.DatabaseUser(),
			env.DatabasePassword(),
			env.DatabaseDBName(),
			env.DatabaseSSLMode(),
		)
		logger.DefaultLogger().Info().
			Timestamp().Str("psqlInfo", psqlInfo).Send()

		config, err := pgxpool.ParseConfig(psqlInfo)
		if err != nil {
			log.Println(err)
		}
		config.MaxConns = 10
		config.MaxConnLifetime = time.Millisecond
		config.MaxConnIdleTime = time.Second
		config.ConnConfig.Logger = zerologadapter.NewLogger(
			zerolog.New(os.Stdout).With().Timestamp().Logger(),
		)

		// connect
		pool, err = pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			err = errors.Wrap(err, "database connect fail")
		}
	})

	return pool, err
}

func Conn(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, errors.Wrap(
			err,
			"pgx connection pool acquire error",
		)
	}
	return conn, nil
}

func ConnCtx(ctx context.Context) (*pgxpool.Conn, error) {
	v := ctx.Value(KEY_DATABASE_CTX)
	if v, ok := v.(*pgxpool.Conn); ok {
		return v, nil
	}
	return nil, errors.New("there is no database connection in the context")
}
