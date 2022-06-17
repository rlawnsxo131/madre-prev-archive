package engine

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
	"github.com/rlawnsxo131/madre-server-v3/external/datastore/rdb"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpmiddleware"
	"github.com/rlawnsxo131/madre-server-v3/external/engine/httpresponse"
	"github.com/rlawnsxo131/madre-server-v3/internal/api"

	"github.com/rlawnsxo131/madre-server-v3/lib/env"
	"github.com/rlawnsxo131/madre-server-v3/lib/logger"
)

const (
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type httpEngine struct {
	db  rdb.Database
	r   *chi.Mux
	srv *http.Server
}

func NewHTTPEngine(db rdb.Database) *httpEngine {
	r := chi.NewRouter()
	e := &httpEngine{
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
	e.RegisterHTTPMiddleware()
	e.RegisterHealthRoute()
	e.RegisterAPIRoute()
	return e
}

func (s *httpEngine) Start() {
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

func (e *httpEngine) RegisterHTTPMiddleware() {
	e.r.Use(chi_middleware.RequestID)
	e.r.Use(chi_middleware.RealIP)
	e.r.Use(httpmiddleware.Logger)
	e.r.Use(httpmiddleware.Recovery)
	e.r.Use(httpmiddleware.AllowHost)
	e.r.Use(httpmiddleware.Cors)
	e.r.Use(httpmiddleware.JWT)
	e.r.Use(httpmiddleware.ContentTypeToJson)
	e.r.Use(chi_middleware.Compress(5))
}

func (e *httpEngine) RegisterHealthRoute() {
	e.r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		rw := httpresponse.NewWriter(w, r)
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

func (e *httpEngine) RegisterAPIRoute() {
	e.r.Route("/api", func(r chi.Router) {
		api.NewAPIV1(r, e.db).Register()
	})
}
