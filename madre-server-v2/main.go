package main

import (
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/lib/env"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/server"
)

func main() {
	db, err := database.DatabaseInstance()
	if err != nil {
		logger.DefaultLogger().Fatal().Timestamp().Err(err).Send()
	}
	defer db.DB.Close()

	if env.IsLocal() {
		database.ExcuteInitSQL(db.DB)
	}

	s := server.NewEngine(db)
	s.Start()
}

func init() {
	logger.DefaultLogger().Info().Timestamp().Msg("init main")
	env.Load()
}
