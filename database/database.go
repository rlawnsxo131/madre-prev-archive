package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/env"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

var (
	instanceDatabase *singletonDatabase
	onceDatabase     sync.Once
)

type Database interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamedGet(id *string, query string, args interface{}) error
}

func DatabaseInstance() (*singletonDatabase, error) {
	var resultError error

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
			Timestamp().Str("database connection info", psqlInfo).Send()

		db, err := sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			resultError = errors.Wrap(err, "sqlx connect fail")
			return
		}

		instanceDatabase = &singletonDatabase{
			DB: db,
			l:  logger.NewBaseLogger(),
		}
		initDatabase(instanceDatabase.DB)
	})

	return instanceDatabase, resultError
}

func initDatabase(db *sqlx.DB) {
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute)
}
