package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/engine"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
)

func main() {
	// db, err := rdb.DatabaseInstance()
	// if err != nil {
	// 	logger.DefaultLogger().Fatal().Timestamp().Err(err).Send()
	// }
	// defer db.DB.Close()

	pool, err := rdb.InitDatabase()
	if err != nil {
		log.Println(err)
	}
	if env.IsLocal() {
		rdb.ExcuteInitSQL(pool)
	}

	e := engine.NewHTTPEngine()
	e.Start()
}

func init() {
	logger.DefaultLogger().Info().Timestamp().Msg("init main")
	env.Load()
}
