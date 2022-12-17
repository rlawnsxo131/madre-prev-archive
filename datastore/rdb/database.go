package rdb

import (
	"context"
	"fmt"
	"os"

	"sync"
	"time"

	pgxlog "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rs/zerolog"
)

var (
	pool         *pgxpool.Pool
	onceDatabase sync.Once
)

func InitDatabasePool() (*pgxpool.Pool, error) {
	var err error

	onceDatabase.Do(func() {
		var config *pgxpool.Config
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			env.DatabaseHost(),
			env.DatabasePort(),
			env.DatabaseUser(),
			env.DatabasePassword(),
			env.DatabaseDBName(),
			env.DatabaseSSLMode(),
		)
		logger.NewDefaultLogger().Add(func(e *zerolog.Event) {
			e.Str("psqlInfo", psqlInfo)
		}).SendInfo()

		config, err = pgxpool.ParseConfig(psqlInfo)
		if err != nil {
			return
		}
		config.MaxConns = 10
		config.MinConns = 0
		config.MaxConnLifetime = time.Minute * 10
		config.MaxConnIdleTime = time.Second * 10

		logLevel := tracelog.LogLevelTrace
		if !env.IsLocal() {
			logLevel = tracelog.LogLevelDebug
		}

		config.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger: pgxlog.NewLogger(
				zerolog.New(os.Stdout).With().Timestamp().Logger(),
			),
			LogLevel: logLevel,
		}

		// connect
		pool, err = pgxpool.NewWithConfig(
			context.Background(),
			config,
		)
		if err != nil {
			err = errors.Wrap(err, "database connect fail")
		}
	})

	return pool, err
}
