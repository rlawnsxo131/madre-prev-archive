package lib

// go get -u github.com/rs/zerolog/log

type httpLogger struct{}

func NewHttpLogger() *httpLogger {
	return &httpLogger{}
}
