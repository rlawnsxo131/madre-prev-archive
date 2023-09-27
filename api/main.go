package main

import (
	"database/sql"
	"log"
	"net/http"
	"runtime"
	"time"

	chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/rlawnsxo131/madre-server/core/httpserver"
	http_middleware "github.com/rlawnsxo131/madre-server/core/httpserver/middleware"
	"github.com/rlawnsxo131/madre-server/core/logger"

	"github.com/rlawnsxo131/madre-server/domain/persistence"

	"github.com/go-sql-driver/mysql"
)

func main() {
	coreCount := runtime.NumCPU()
	runtime.GOMAXPROCS(coreCount - 1)

	log.Printf(
		"core count: %v, current max use cpu count: %v",
		coreCount,
		runtime.GOMAXPROCS(0),
	)

	// @TODO 임시 테스트용 db connection
	// db, err := sql.Open("mysql", "user:password@/dbname")
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "1234",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		Collation:            "utf8mb4_0900_ai_ci",
		Loc:                  time.Local,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               "madre",
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}

	db := sql.OpenDB(connector)
	err = persistence.ExcuteInitSQL(db)
	if err != nil {
		log.Fatal(err)
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

	s.Router().Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// @TODO dev 에서만 사용하도록 처리
	s.Router().Mount("/debug", chi_middleware.Profiler())

	s.Start()

}
