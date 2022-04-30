package database

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rs/zerolog"
)

// TODO: Thinking about redefining the database and working on it

type Database interface {
	Begin()
	Commit()
	Rollback()
}

var (
	once             sync.Once
	instanceDatabase *singletonDatabase
)

func GetDatabaseInstance() (Database, error) {
	var err error

	once.Do(func() {
		l := zerolog.New(os.Stderr)
		instanceDatabase = &singletonDatabase{
			logger: &l,
		}

		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.NewDefaultLogger().Logger.Info().Str("database connection info", psqlInfo).Send()

		instanceDatabase.Database, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			return
		}

		instanceDatabase.Database.SetMaxIdleConns(5)
		instanceDatabase.Database.SetMaxOpenConns(5)
		instanceDatabase.Database.SetConnMaxLifetime(time.Minute)
	})

	return instanceDatabase, err
}
