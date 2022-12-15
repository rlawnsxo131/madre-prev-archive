package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/engine"
	"github.com/rlawnsxo131/madre-server-v3/core/engine/logger"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rs/zerolog"
)

func main() {
	pool, err := rdb.InitDatabasePool()
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
	logger.DefaultLogger().Add(func(e *zerolog.Event) {
		e.Str("message", "init main")
	}).SendInfo()

	env.Load()
}
