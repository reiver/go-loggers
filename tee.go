package loggers


type Tee struct {
	Logger1 Logger
	Logger2 Logger
}


func (logger Tee) Debug(v ...interface{}) {
	logger.Logger1.Debug(v...)
	logger.Logger2.Debug(v...)
}

func (logger Tee) Debugf(format string, v ...interface{}) {
	logger.Logger1.Debugf(format, v...)
	logger.Logger2.Debugf(format, v...)
}


func (logger Tee) Error(v ...interface{}) {
	logger.Logger1.Error(v...)
	logger.Logger2.Error(v...)
}

func (logger Tee) Errorf(format string, v ...interface{}) {
	logger.Logger1.Errorf(format, v...)
	logger.Logger2.Errorf(format, v...)
}


func (logger Tee) Print(v ...interface{}) {
	logger.Logger1.Print(v...)
	logger.Logger2.Print(v...)
}

func (logger Tee) Printf(format string, v ...interface{}) {
	logger.Logger1.Printf(format, v...)
	logger.Logger2.Printf(format, v...)
}


func (logger Tee) Trace(v ...interface{}) {
	logger.Logger1.Trace(v...)
	logger.Logger2.Trace(v...)
}

func (logger Tee) Tracef(format string, v ...interface{}) {
	logger.Logger1.Tracef(format, v...)
	logger.Logger2.Tracef(format, v...)
}


func (logger Tee) Warn(v ...interface{}) {
	logger.Logger1.Warn(v...)
	logger.Logger2.Warn(v...)
}

func (logger Tee) Warnf(format string, v ...interface{}) {
	logger.Logger1.Warnf(format, v...)
	logger.Logger2.Warnf(format, v...)
}
