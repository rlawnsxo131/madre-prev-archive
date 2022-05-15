package main

import (
	"log"

	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/env"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db, err := database.DatabaseInstance()
	if err != nil {
		logger.DefaultLogger().Fatal().Timestamp().Err(err).Msg("")
	}
	defer db.DB.Close()

	if env.IsLocal() {
		log.Println("local")
		database.ExcuteInitSQL(db.DB)
	}

	s := server.New(db)
	s.Start()
}

func init() {
	log.Println("init main")
	env.Load()
}
