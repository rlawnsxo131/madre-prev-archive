package database

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rs/zerolog"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "madre"
	password = "1234"
	dbname   = "madre"
)

var (
	once             sync.Once
	instanceDatabase *singletonDatabase
)

func GetDatabaseInstance() (*singletonDatabase, error) {
	var err error

	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.NewDefaultLogger().
			Info().Str("database connection info", psqlInfo).Send()

		db, err := sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			err = errors.Wrap(err, "sqlx: connect fail")
			return
		}

		l := zerolog.New(os.Stderr)
		instanceDatabase = &singletonDatabase{
			DB:     db,
			logger: &l,
		}
		initDatabase(instanceDatabase.DB)
	})

	return instanceDatabase, err
}

func initDatabase(db *sqlx.DB) {
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute)
}
