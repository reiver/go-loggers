package loggers

// Discard is a Logger on which all log calls succeed without doing anything.
//
// The logging data is essentially "thrown away".
var Discard Logger

func init() {
	Discard = discardLogger{}
}

type discardLogger struct{}

func (discardLogger) Debug(v ...interface{}) {}
func (discardLogger) Debugf(format string, v ...interface{}) {}

func (discardLogger) Error(v ...interface{}) {}
func (discardLogger) Errorf(format string, v ...interface{}) {}

func (discardLogger) Fatal(v ...interface{}) {}
func (discardLogger) Fatalf(format string, v ...interface{}) {}

func (discardLogger) Panic(v ...interface{}) {}
func (discardLogger) Panicf(format string, v ...interface{}) {}

func (discardLogger) Print(v ...interface{}) {}
func (discardLogger) Printf(format string, v ...interface{}) {}

func (discardLogger) Trace(v ...interface{}) {}
func (discardLogger) Tracef(format string, v ...interface{}) {}

func (discardLogger) Warn(v ...interface{}) {}
func (discardLogger) Warnf(format string, v ...interface{}) {}
