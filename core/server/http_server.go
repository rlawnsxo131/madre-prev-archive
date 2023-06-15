package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server-v3/core/env"
	"github.com/rlawnsxo131/madre-server-v3/core/logger"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httplogger"
	"github.com/rlawnsxo131/madre-server-v3/core/server/httpmiddleware"
	"github.com/rs/zerolog"
)

const (
	WRITE_TIMEOUT = time.Second * 15
	READ_TIMEOUT  = time.Second * 15
	IDLE_TIMEOUT  = time.Second * 60
)

type httpServer struct {
	r   *chi.Mux
	srv *http.Server
}

func NewHTTPServer() *httpServer {
	r := chi.NewRouter()
	s := &httpServer{
		r: r,
		srv: &http.Server{
			Addr: "0.0.0.0:" + env.Port(),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: WRITE_TIMEOUT,
			ReadTimeout:  READ_TIMEOUT,
			IdleTimeout:  IDLE_TIMEOUT,
			Handler:      r,
		},
	}
	return s
}

func (s *httpServer) Start() {
	le := logger.DefaultLogger.NewLogEntry()

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
				le.Add(func(e *zerolog.Event) {
					e.Str("message", "graceful shutdown timed out.. forcing exit.")
				}).SendFatal()
			}
		}()

		// Trigger graceful shutdown
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {
			le.Add(func(e *zerolog.Event) {
				e.Err(err)
			}).SendFatal()
		}
		srvCtxCancel()
	}()

	// Run the server
	le.Add(func(e *zerolog.Event) {
		e.Str("message", "going to listen on port "+env.Port())
	}).SendInfo()
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		le.Add(func(e *zerolog.Event) {
			e.Err(err)
		}).SendFatal()
	}

	// Wait for server context to be stopped
	<-srvCtx.Done()
}

func (s *httpServer) RegisterHTTPMiddleware(hl httplogger.HTTPLogger) {
	s.r.Use(chi_middleware.RequestID)
	s.r.Use(chi_middleware.RealIP)
	s.r.Use(httpmiddleware.Logger(hl))
	s.r.Use(httpmiddleware.Recovery)
	s.r.Use(httpmiddleware.AllowHost)
	s.r.Use(httpmiddleware.Cors)
	s.r.Use(httpmiddleware.JWT)
	s.r.Use(chi_middleware.Compress(5))
}

func (s *httpServer) Route() *chi.Mux {
	return s.r
}
