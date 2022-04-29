package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
)

// TODO: Thinking about redefining the database and working on it

var (
	once             sync.Once
	instanceDatabase *singletonDatabse
)

type Database interface{}

type singletonDatabse struct {
	Database *sqlx.DB
}

func GetDatabaseInstance() (*singletonDatabse, error) {
	var err error

	once.Do(func() {
		instanceDatabase = &singletonDatabse{}

		psqlInfo := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.Logger.Info().Str("database connection info", psqlInfo).Send()

		instanceDatabase.Database, err = sqlx.Connect("postgres", psqlInfo)
		if err != nil {
			return
		}

		initDatabase(instanceDatabase.Database)
	})

	return instanceDatabase, err
}

func initDatabase(db *sqlx.DB) {
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute)
}
