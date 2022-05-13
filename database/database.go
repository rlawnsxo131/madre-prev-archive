package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "madre"
	password = "1234"
	dbname   = "madre"
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

func GetDatabaseInstance() (*singletonDatabase, error) {
	var resultError error

	onceDatabase.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.GetDefaultLogger().
			Info().Str("database connection info", psqlInfo).Msg("")

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
