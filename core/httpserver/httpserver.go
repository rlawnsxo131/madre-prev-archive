package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
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

			}
		}()

		// Trigger graceful shutdown
		err := s.srv.Shutdown(shutdownCtx)
		if err != nil {

		}
		srvCtxCancel()
	}()

	// Run the server
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {

	}

	// Wait for server context to be stopped
	<-srvCtx.Done()
}

// func (s *httpServer) RegisterHTTPMiddleware(hl httplogger.HTTPLogger) {
// 	s.r.Use(chi_middleware.RequestID)
// 	s.r.Use(chi_middleware.RealIP)
// 	s.r.Use(httpmiddleware.Logger(hl))
// 	s.r.Use(httpmiddleware.Recovery)
// 	s.r.Use(httpmiddleware.AllowHost)
// 	s.r.Use(httpmiddleware.Cors)
// 	s.r.Use(httpmiddleware.JWT)
// 	s.r.Use(chi_middleware.Compress(5))
// }

func (s *httpServer) Route() *chi.Mux {
	return s.r
}
