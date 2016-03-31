package loggers


import (
	"github.com/reiver/go-oi"

	"fmt"
	"io"
	"os"
)


type Writers struct {
	FatalWriter io.Writer
	PanicWriter io.Writer
	ErrorWriter io.Writer
	WarnWriter  io.Writer
	PrintWriter io.Writer
	DebugWriter io.Writer
	TraceWriter io.Writer
}


func (logger Writers) Debug(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.DebugWriter, msg)
}

func (logger Writers) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.DebugWriter, msg)
}


func (logger Writers) Error(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.ErrorWriter, msg)
}

func (logger Writers) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.ErrorWriter, msg)
}


func (logger Writers) Fatal(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.FatalWriter, msg)
	os.Exit(1)
}

func (logger Writers) Fatalf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.FatalWriter, msg)
	os.Exit(1)
}


func (logger Writers) Panic(v ...interface{}) {
	err := fmt.Errorf("%s", fmt.Sprint(v...))
	oi.LongWriteString(logger.PanicWriter, err.Error())
	panic(err)
}

func (logger Writers) Panicf(format string, v ...interface{}) {
	err := fmt.Errorf(format, v...)
	oi.LongWriteString(logger.PanicWriter, err.Error())
	panic(err)
}


func (logger Writers) Print(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.PrintWriter, msg)
}

func (logger Writers) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.PrintWriter, msg)
}


func (logger Writers) Trace(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.TraceWriter, msg)
}

func (logger Writers) Tracef(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.TraceWriter, msg)
}


func (logger Writers) Warn(v ...interface{}) {
	msg := fmt.Sprint(v...)
	oi.LongWriteString(logger.WarnWriter, msg)
}

func (logger Writers) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	oi.LongWriteString(logger.WarnWriter, msg)
}
