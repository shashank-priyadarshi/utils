package ports

type Logger interface {
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(error, ...interface{})
	Fatal(error, ...interface{})
	Debug(string, ...interface{})
	With(args ...interface{})
	/* TODO
	   Enable logging to file
	   Enable log rotation
	   Enable zipping log files
	*/
}
