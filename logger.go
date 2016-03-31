package loggers


// Logger represents a logger with levels.
//
// Logger includes the Panic, Panicf, Print, Printf, Trace, and Tracef methods from the *log.Logger type in the Go standard library.
//
// But it also adds the Debug, Debugf, Error, Errorf, Trace, Tracef, Warn, and Warnf methods that are common with loggers with "levels".
//
// The "levels" in terms of least verbose to most verbose are:
//
//	· Fatal
//	· Panic
//	· Error
//	· Warn
//	· Print
//	· Debug
//	· Trace
//
type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})

	Print(...interface{})
	Printf(string, ...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
}
