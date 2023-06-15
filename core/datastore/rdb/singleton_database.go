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
	"github.com/rlawnsxo131/madre-server-v3/core/env"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
	"github.com/rs/zerolog"
)

var (
	instanceDatabase *singletonDatabase
	onceDatabase     sync.Once
)

type SingletonDatabase interface {
	Pool() *pgxpool.Pool
	ClosePool()
	Conn() (*pgxpool.Conn, error)
}

type singletonDatabase struct {
	pool *pgxpool.Pool
}

func DBInstance() (SingletonDatabase, error) {
	var err error

	onceDatabase.Do(func() {
		pool, initPoolErr := initDatabasePool()
		if err != nil {
			err = initPoolErr
			return
		}

		instanceDatabase = &singletonDatabase{
			pool: pool,
		}
	})

	return instanceDatabase, err
}

func (sd *singletonDatabase) Pool() *pgxpool.Pool {
	return sd.pool
}

func (sd *singletonDatabase) ClosePool() {
	sd.pool.Close()
}

func (sd *singletonDatabase) Conn() (*pgxpool.Conn, error) {
	conn, err := sd.pool.Acquire(context.Background())
	if err != nil {
		return nil, errors.Wrap(
			err,
			"pgx connection pool acquire error",
		)
	}
	return conn, nil
}

// initialize pgx pool
func initDatabasePool() (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.DatabaseHost(),
		env.DatabasePort(),
		env.DatabaseUser(),
		env.DatabasePassword(),
		env.DatabaseDBName(),
		env.DatabaseSSLMode(),
	)
	logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
		e.Str("psqlInfo", psqlInfo)
	}).SendInfo()

	config, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "initDatabasePool ParseConfig error")
	}
	config.MaxConns = 10
	config.MinConns = 0
	config.MaxConnLifetime = time.Minute * 10
	config.MaxConnIdleTime = time.Second * 10

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logLevel := tracelog.LogLevelDebug
	if env.IsLocal() {
		logLevel = tracelog.LogLevelTrace
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: pgxlog.NewLogger(
			logger,
		),
		LogLevel: logLevel,
	}

	// connect
	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		config,
	)
	if err != nil {
		return nil, errors.Wrap(err, "initDatabasePool NewWithConfig error")
	}

	return pool, nil
}
