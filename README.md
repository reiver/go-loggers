# go-loggers

Package **loggers** provides useful tools for dealing with loggers, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-loggers

[![GoDoc](https://godoc.org/github.com/reiver/go-loggers?status.svg)](https://godoc.org/github.com/reiver/go-loggers)


## Example
```
logger := loggers.Writers{
	FatalWriter: os.Stderr,
	PanicWriter: os.Stderr,
	ErrorWriter: os.Stderr,
	WarnWriter:  os.Stderr,
	PrintWriter: os.Stdout,
	DebugWriter: os.Stdout,
	TraceWriter: osutil.Discard,
}

logger.Print("Hello world!")

logger.Warn("Be very very careful!")

logger.Fatalf("Finish %s", name)
``
