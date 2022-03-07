package lib

// go get -u github.com/rs/zerolog/log

var logger *httpLogger

type HttpLogger interface{}

type httpLogger struct{}

func NewHttpLogger() *httpLogger {
	if logger == nil {
		logger = &httpLogger{}
	}
	return logger
}
