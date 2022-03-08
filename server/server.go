package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rlawnsxo131/madre-server-v2/domain/auth"
	"github.com/rlawnsxo131/madre-server-v2/domain/data"
	"github.com/rlawnsxo131/madre-server-v2/domain/user"
	"github.com/rlawnsxo131/madre-server-v2/middleware"
	"github.com/urfave/negroni"
)

const (
	port         = "5000"
	writeTimeout = time.Second * 15
	readTimeout  = time.Second * 15
	idleTimeout  = time.Second * 60
)

type server struct {
	router     *mux.Router
	rootRouter *mux.Router
	handler    *negroni.Negroni
	httpServer *http.Server
	db         *sqlx.DB
}

func New(db *sqlx.DB) *server {
	s := &server{
		router:     mux.NewRouter(),
		rootRouter: mux.NewRouter().PathPrefix("/").Subrouter(),
		handler:    negroni.New(negroni.NewRecovery()),
		db:         db,
	}
	s.setupRouteAndMiddleware()
	s.connectHandlerAndRouter()
	s.setupHttpServer()
	return s
}

func (s *server) Start() {
	log.Println("Going to listen on port", port)
	log.Fatal(s.httpServer.ListenAndServe())
}

func (s *server) setupRouteAndMiddleware() {
	apiRouter := s.rootRouter.NewRoute().PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Method":  r.Method,
			"Host":    r.Host,
			"Path":    r.URL.Path,
			"Referer": r.Header.Get("Referer"),
			"Cookies": fmt.Sprint(r.Cookies()),
		}
		json.NewEncoder(rw).Encode(data)
	})

	s.router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(middleware.CorsMiddleware),
		negroni.HandlerFunc(middleware.SetDatabaseContextMiddleware(s.db)),
		negroni.HandlerFunc(middleware.SetResponseContentTypeMiddleware),
		negroni.Wrap(apiRouter),
	))

	v1 := apiRouter.NewRoute().PathPrefix("/v1").Subrouter()
	auth.SetupRoute(v1)
	user.SetupRoute(v1)
	data.SetupRoute(v1)
}

func (s *server) connectHandlerAndRouter() {
	s.handler.UseHandler(s.router)
}

func (s *server) setupHttpServer() {
	s.httpServer = &http.Server{
		Addr: "0.0.0.0:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
		Handler:      s.handler,
	}
}
