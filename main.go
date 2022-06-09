package main

import (
	"github.com/rlawnsxo131/madre-server-v3/internal/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/internal/engine"
	"github.com/rlawnsxo131/madre-server-v3/internal/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/internal/lib/logger"
)

func main() {
	db, err := rdb.DatabaseInstance()
	if err != nil {
		logger.DefaultLogger().Fatal().Timestamp().Err(err).Send()
	}
	defer db.DB.Close()

	if env.IsLocal() {
		rdb.ExcuteInitSQL(db.DB)
	}

	e := engine.NewHTTPEngine(db)
	e.Start()
}

func init() {
	logger.DefaultLogger().Info().Timestamp().Msg("init main")
	env.Load()
}
