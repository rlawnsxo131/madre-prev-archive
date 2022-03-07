package lib

// go get -u github.com/rs/zerolog/log

var logger *httpLogger

type HttpLogger interface{}

type httpLogger struct{}

func GetHttpLogger() *httpLogger {
	once.Do(func() {
		logger = &httpLogger{}
	})
	return logger
}
