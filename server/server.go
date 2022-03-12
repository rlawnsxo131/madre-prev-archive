package server

import (
	"fmt"
	"log"
	"net/http"
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
	log.Println("Going to listen on port", port)
	log.Fatal(s.httpServer.ListenAndServe())
}

func (s *server) applyHealthSettings() {
	s.router.Use(
		middleware.HttpLogger,
		middleware.Recovery,
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
		middleware.Cors,
		middleware.SetHttpContextValues(s.db),
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
