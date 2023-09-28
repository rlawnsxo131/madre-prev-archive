package main

import (
	"context"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"

	routerv1 "github.com/rlawnsxo131/madre-server/api/router/v1"
	"github.com/rlawnsxo131/madre-server/core/funk"
	"github.com/rlawnsxo131/madre-server/core/httpserver"
	http_middleware "github.com/rlawnsxo131/madre-server/core/httpserver/middleware"

	"github.com/rlawnsxo131/madre-server/core/logger"
	"github.com/rlawnsxo131/madre-server/core/rdb"

	"github.com/rlawnsxo131/madre-server/domain/persistence"
)

func main() {
	coreCount := runtime.NumCPU()
	runtime.GOMAXPROCS(coreCount - 1)

	logger.DefaultLogger().
		Info().
		Int("core count", coreCount).
		Int("max use cpu count", runtime.GOMAXPROCS(0)).Send()

	db, err := rdb.CreateConnection(&mysql.Config{
		User:                 "root",
		Passwd:               "1234",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		Collation:            "utf8mb4_0900_ai_ci",
		Loc:                  time.Local,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		ParseTime:            true,
		DBName:               "madre",
	})
	if err != nil {
		logger.DefaultLogger().Fatal().Err(err).Send()
	}

	if os.Getenv("APP_ENV") == "local" {
		tx, err := db.BeginTx(context.Background(), nil)
		if err != nil {
			logger.DefaultLogger().Fatal().Err(err).Send()
		}

		err = persistence.ExcuteInitSQL(tx)
		if err != nil {
			tx.Rollback()
			logger.DefaultLogger().Fatal().Err(err).Send()
		}
		tx.Commit()
	}

	s := httpserver.New("127.0.0.1:5001")

	s.Use(chi_middleware.RequestID)
	s.Use(chi_middleware.RealIP)
	s.Use(logger.HTTPMiddleware(logger.DefaultHTTPLogger))
	s.Use(http_middleware.Recover)
	s.Use(
		http_middleware.AllowHost(
			[]string{"localhost:8080", "localhost:5001"},
		),
	)
	s.Use(http_middleware.Cors(
		[]string{"http://localhost:8080", "http://localhost:5001"},
	))
	// s.Use(http_middleware.JWT)
	s.Use(chi_middleware.Compress(5))

	// register routes
	s.Router().Route("/", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})

		r.Route("/api", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				routerv1.NewAuthRouter(r)
			})
		})
	})

	if funk.Contains[string]([]string{"local", "dev"}, os.Getenv("APP_ENV")) {
		s.Router().Mount("/debug", chi_middleware.Profiler())
	}

	s.Start()
}
