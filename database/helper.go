package database

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v2/constants"
	"github.com/rlawnsxo131/madre-server-v2/lib/syncmap"
)

func ExcuteInitSQL(db *sqlx.DB) {
	file, err := ioutil.ReadFile("./database/init.sql")
	if err != nil {
		panic(err)
	}

	queries := strings.Split(string(file), "\n\n")
	for _, query := range queries {
		db.MustExec(query)
	}
}

func LoadDBFromHttpSyncMapContext(ctx context.Context) (*singletonDatabase, error) {
	syncMap, err := syncmap.GetFromHttpContext(ctx)
	if err != nil {
		return nil, err
	}

	if db, ok := syncMap.Load(constants.Key_HttpContextDB); ok {
		if db, ok := db.(*singletonDatabase); ok {
			return db, nil
		}
	}

	return nil, errors.New("DB is not exist")
}
