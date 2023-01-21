package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v3/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
	"github.com/rlawnsxo131/madre-server-v3/server"
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

	s := server.NewHTTPServer()
	s.Start()
}

func init() {
	logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
		e.Str("message", "init main")
	}).SendInfo()

	env.Load()
}
