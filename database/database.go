package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rs/zerolog"
)

// TODO: Thinking about redefining the database and working on it

type Database interface {
	Begin()
	Commit()
	Rollback()
	initDatabase()
}

var (
	once             sync.Once
	instanceDatabase *singletonDatabase
)

func GetDatabaseInstance() (Database, error) {
	var err error

	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.NewDefaultLogger().Logger.Info().Str("database connection info", psqlInfo).Send()

		var database *sqlx.DB
		database, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			return
		}

		l := zerolog.New(os.Stderr)
		instanceDatabase = &singletonDatabase{
			database: database,
			logger:   &l,
		}
		instanceDatabase.initDatabase()
	})

	return instanceDatabase, err
}
