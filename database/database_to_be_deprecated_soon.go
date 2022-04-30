package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "madre"
	password = "1234"
	dbname   = "madre"
)

var (
	sqlxDB *sqlx.DB
)

func GetDatabase() (*sqlx.DB, error) {
	var err error

	once.Do(func() {
		// get database connection cofig
		dbString := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
		logger.NewDefaultLogger().
			Logger.
			Info().
			Str("database connection info", dbString).
			Send()

		// database connect
		sqlxDB, err = sqlx.Connect("postgres", dbString)
		if err != nil {
			err = errors.Wrap(err, "sqlx: connect fail")
			return
		}

		// initialize database config
		sqlxDB.SetMaxIdleConns(5)
		sqlxDB.SetMaxOpenConns(5)
		sqlxDB.SetConnMaxLifetime(time.Minute)
	})

	return sqlxDB, err
}

func GetDatabseFromHttpContext(ctx context.Context) (*sqlx.DB, error) {
	syncMap, err := syncmap.GetFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}

	if sqlxDB, ok := syncMap.Load(constants.Key_HttpContextDB); ok {
		if sqlxDB, ok := sqlxDB.(*sqlx.DB); ok {
			return sqlxDB, nil
		}
	}

	return nil, errors.New("DB is not exist")
}

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		panic(err)
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		sqlxDB.MustExec(query)
	}
}
