package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v3/core"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rs/zerolog"
)

func main() {
	pool, err := rdb.InitDatabasePool()
	if err != nil {
		log.Println(err)
	}
	defer pool.Close()

	if env.IsLocal() {
		rdb.ExcuteInitSQL(pool)
	}

	e := core.NewHTTPServer()
	e.Start()
}

func init() {
	logger.NewDefaultLogger().Add(func(e *zerolog.Event) {
		e.Str("message", "init main")
	}).SendInfo()

	env.Load()
}
