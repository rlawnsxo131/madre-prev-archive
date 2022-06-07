package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/env"
	"github.com/rlawnsxo131/madre-server-v2/lib/logger"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/middleware"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type engine struct {
	db  database.Database
	r   *chi.Mux
	srv *http.Server
}

func New(db database.Database) *engine {
	r := chi.NewRouter()
	e := &engine{
		db: db,
		r:  r,
		srv: &http.Server{
			Addr: "0.0.0.0:" + env.Port(),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
			IdleTimeout:  idleTimeout,
			Handler:      r,
		},
	}
	e.RegisterMiddleware()
	e.RegisterHealthRoute()
	e.RegisterAPIRoutes()
	return e
}

func (s *engine) Start() {
	// Server run context
	srvCtx, srvCtxCancel := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCtxCancel := context.WithTimeout(srvCtx, 30*time.Second)
		defer shutdownCtxCancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.DefaultLogger().Fatal().
					Timestamp().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {
			logger.DefaultLogger().Fatal().Timestamp().Err(err).Send()
		}
		srvCtxCancel()
	}()

	// Run the server
	logger.DefaultLogger().Info().
		Timestamp().Msg("going to listen on port " + env.Port())
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.DefaultLogger().Fatal().Timestamp().Err(err).Send()
	}

	// Wait for server context to be stopped
	<-srvCtx.Done()
}

func (e *engine) RegisterMiddleware() {
	e.r.Use(chi_middleware.RequestID)
	e.r.Use(chi_middleware.RealIP)
	e.r.Use(middleware.HTTPLogger)
	e.r.Use(middleware.Recovery)
	e.r.Use(middleware.AllowHost)
	e.r.Use(middleware.Cors)
	e.r.Use(middleware.JWT)
	e.r.Use(middleware.ContentTypeToJson)
	e.r.Use(chi_middleware.Compress(5))
}

func (e *engine) RegisterHealthRoute() {
	e.r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		data := map[string]string{
			"Method":  r.Method,
			"Host":    r.Host,
			"Path":    r.URL.Path,
			"Referer": r.Header.Get("Referer"),
			"Cookies": fmt.Sprint(r.Cookies()),
		}
		rw.Write(data)
	})
}

func (e *engine) RegisterAPIRoutes() {
	e.r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			auth.RegisterRoutes(r, e.db)
			user.RegisterRoutes(r, e.db)
		})
	})
}
