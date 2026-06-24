package ecs

// any logger that has the methods can be plugged in for use by
// the game engine
type Logger interface {
	Info(msg string)
	Warn(msg string)
	Err(err error)
}

// noop logger serves as a placeholder and supplies the log methods
// in the event a logger is not set -- nothing is actually
// logged with the default logger
type defaultLogger struct{}

func (defaultLogger) Info(string) {}
func (defaultLogger) Warn(string) {}
func (defaultLogger) Err(error)   {}
