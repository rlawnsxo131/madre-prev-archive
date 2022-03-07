package lib

// go get -u github.com/rs/zerolog/log

type HttpLogger interface{}

type httpLogger struct{}

var logger *httpLogger

func NewHttpLogger() *httpLogger {
	if logger == nil {
		logger = &httpLogger{}
	}
	return logger
}
