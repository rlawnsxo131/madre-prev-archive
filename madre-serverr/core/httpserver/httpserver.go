package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rlawnsxo131/madre-server/core/of"
)

const (
	_writeTimeout = time.Second * 15
	_readTimeout  = time.Second * 15
	_idleTimeout  = time.Second * 60
)

type httpServer struct {
	r   *chi.Mux
	srv *http.Server
}

func New(addr string) *httpServer {
	r := chi.NewRouter()
	s := &httpServer{
		r: r,
		srv: &http.Server{
			Addr: addr,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: _writeTimeout,
			ReadTimeout:  _readTimeout,
			IdleTimeout:  _idleTimeout,
			Handler:      r,
		},
	}
	return s
}

func (s *httpServer) Start() {
	// Server run context
	srvCtx, srvStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownStopCtx := context.WithTimeout(srvCtx, 30*time.Second)
		defer shutdownStopCtx()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				of.DefaultLogger().Fatal().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {
			of.DefaultLogger().Fatal().Err(err).Send()
		}
		srvStopCtx()
	}()

	// Run the server
	of.DefaultLogger().Info().Msgf("server start at %v", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		of.DefaultLogger().Fatal().Err(err).Send()
	}

	// Wait for server context to be stopped
	<-srvCtx.Done()
}

func (s *httpServer) Use(middleware func(http.Handler) http.Handler) {
	s.r.Use(middleware)
}

func (s *httpServer) Router() *chi.Mux {
	return s.r
}
