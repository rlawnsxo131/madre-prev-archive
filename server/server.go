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
	"github.com/rlawnsxo131/madre-server-v2/database"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/domain/data"
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
	db         database.Database
	router     *mux.Router
	httpServer *http.Server
}

func New(db database.Database) *server {
	s := &server{
		db:     db,
		router: mux.NewRouter(),
	}
	s.applyMiddleware()
	s.applyHealthRoute()
	s.applyApiRoutes()
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
		log.Println("Going to listen on port", port)
		err := s.httpServer.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
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

func (s *server) applyMiddleware() {
	s.router.Use(
		middleware.HTTPLogger,
		middleware.Recovery,
		middleware.AllowHost,
		middleware.Cors,
		middleware.JWT,
		middleware.ContentTypeToJson,
	)
}

func (s *server) applyHealthRoute() {
	s.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		rw := response.NewWriter(w, r)
		data := map[string]string{
			"Method":  r.Method,
			"Host":    r.Host,
			"Path":    r.URL.Path,
			"Referer": r.Header.Get("Referer"),
			"Cookies": fmt.Sprint(r.Cookies()),
		}
		rw.Compress(data)
	})
}

func (s *server) applyApiRoutes() {
	api := s.router.NewRoute().PathPrefix("/api").Subrouter()
	v1 := api.NewRoute().PathPrefix("/v1").Subrouter()

	auth.RegisterController(v1, s.db)
	user.RegisterController(v1, s.db)
	data.RegisterController(v1, s.db)
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
