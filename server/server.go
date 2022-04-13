package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/domain/data"
	"github.com/rlawnsxo131/madre-server-v2/domain/temp"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/lib/response"
	"github.com/rlawnsxo131/madre-server-v2/middleware"
)

const (
	port         = "5000"
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type server struct {
	router     *mux.Router
	httpServer *http.Server
	db         *sqlx.DB
}

func New(db *sqlx.DB) *server {
	s := &server{
		router: mux.NewRouter(),
		db:     db,
	}
	s.applyHealthSettings()
	s.applyRoutesAndMiddlewares()
	s.applyHttpServer()
	return s
}

func (s *server) Start() {
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)
	flag.Parse()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
		log.Println("Going to listen on port", port)
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.httpServer.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func (s *server) applyHealthSettings() {
	s.router.Use(
		middleware.HttpLogger,
		middleware.Recovery,
		middleware.Cors,
		middleware.SetSyncMapContext,
		middleware.JWT,
	)
	s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writer := response.NewHttpWriter(w, r)
		data := map[string]string{
			"Method":  r.Method,
			"Host":    r.Host,
			"Path":    r.URL.Path,
			"Referer": r.Header.Get("Referer"),
			"Cookies": fmt.Sprint(r.Cookies()),
		}
		writer.WriteCompress(data)
	})
}

func (s *server) applyRoutesAndMiddlewares() {
	api := s.router.NewRoute().PathPrefix("/api").Subrouter()
	api.Use(
		middleware.SetDBContext(s.db),
		middleware.SetResponseContentTypeJson,
	)
	v1 := api.NewRoute().PathPrefix("/v1").Subrouter()
	auth.ApplyRoutes(v1)
	user.ApplyRoutes(v1)
	data.ApplyRoutes(v1)
	temp.ApplyRoutes(v1)
}

func (s *server) applyHttpServer() {
	s.httpServer = &http.Server{
		Addr: "0.0.0.0:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      s.router,
	}
}
