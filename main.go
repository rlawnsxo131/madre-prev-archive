package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server-v3/core/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/core/env"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
	"github.com/rlawnsxo131/madre-server-v3/core/server"
	api "github.com/rlawnsxo131/madre-server-v3/internal/api"
	apiv1 "github.com/rlawnsxo131/madre-server-v3/internal/api/v1"
	"github.com/rs/zerolog"
)

func main() {
	db, err := rdb.DBInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer db.ClosePool()

	if env.IsLocal() {
		rdb.ExcuteInitSQL(db)
	}

	s := server.NewHTTPServer()
	s.RegisterHTTPMiddleware(db)

	// routes
	s.Route().Route("/", func(r chi.Router) {
		// health
		api.NewHealthRoute().Register(r)

		// api
		r.Route("/api", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				apiv1.NewAuthRoute().Register(r)
				apiv1.NewMeRoute().Register(r)
			})
		})
	})

	s.Start()
}

func init() {
	logger.DefaultLogger.NewLogEntry().Add(func(e *zerolog.Event) {
		e.Str("message", "init main")
	}).SendInfo()

	env.Load()
}
