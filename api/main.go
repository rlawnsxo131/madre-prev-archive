package main

import (
	"log"
	"net/http"
	"runtime"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/rlawnsxo131/madre-server/core/httpserver"
	"github.com/rlawnsxo131/madre-server/core/httpserver/httpmiddleware"
)

func main() {
	coreCount := runtime.NumCPU()
	runtime.GOMAXPROCS(coreCount - 1)

	log.Printf(
		"core count: %v, current max use cpu count: %v",
		coreCount,
		runtime.GOMAXPROCS(0),
	)

	s := httpserver.New("0.0.0.0:5001")

	s.Use(chi_middleware.RequestID)
	s.Use(chi_middleware.RealIP)
	// s.Use(httpmiddleware.Logger(hl))
	s.Use(httpmiddleware.Recover)
	s.Use(
		httpmiddleware.AllowHost(
			[]string{"localhost:8080", "localhost:5001"},
		),
	)
	s.Use(httpmiddleware.Cors(
		[]string{"http://localhost:8080", "http://localhost:5001"},
	))
	// s.r.Use(httpmiddleware.JWT)
	s.Use(chi_middleware.Compress(5))

	s.Route().Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// @TODO dev 에서만 사용하도록 처리
	// mux 가 모두 붙은 이후 사용해야 한다
	s.Route().Mount("/debug", chi_middleware.Profiler())

	s.Start()
}
