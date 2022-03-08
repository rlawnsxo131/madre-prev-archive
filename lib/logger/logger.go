package logger

type Logger interface{}
type logger struct{}

func NewLogger() Logger {
	return &logger{}
}
