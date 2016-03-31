package loggers


import (
	"github.com/reiver/go-oi"

	"bytes"
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
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.DebugWriter, buffer.Bytes())
}

func (logger Writers) Debugf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.DebugWriter, buffer.Bytes())
}


func (logger Writers) Error(v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.ErrorWriter, buffer.Bytes())
}

func (logger Writers) Errorf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.ErrorWriter, buffer.Bytes())
}


func (logger Writers) Fatal(v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.FatalWriter, buffer.Bytes())
	os.Exit(1)
}

func (logger Writers) Fatalf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.FatalWriter, buffer.Bytes())
	os.Exit(1)
}


func (logger Writers) Panic(v ...interface{}) {
	var buffer bytes.Buffer

	err := fmt.Errorf("%s", fmt.Sprint(v...))

	buffer.WriteString(err.Error())
	buffer.WriteByte('\n')

	oi.LongWrite(logger.PanicWriter, buffer.Bytes())
	panic(err)
}

func (logger Writers) Panicf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	err := fmt.Errorf(format, v...)

	buffer.WriteString(err.Error())
	buffer.WriteByte('\n')

	oi.LongWrite(logger.PanicWriter, buffer.Bytes())
	panic(err)
}


func (logger Writers) Print(v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.PrintWriter, buffer.Bytes())
}

func (logger Writers) Printf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.PrintWriter, buffer.Bytes())
}


func (logger Writers) Trace(v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.TraceWriter, buffer.Bytes())
}

func (logger Writers) Tracef(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.TraceWriter, buffer.Bytes())
}


func (logger Writers) Warn(v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprint(&buffer, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.WarnWriter, buffer.Bytes())
}

func (logger Writers) Warnf(format string, v ...interface{}) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, format, v...)
	buffer.WriteByte('\n')

	oi.LongWrite(logger.WarnWriter, buffer.Bytes())
}
