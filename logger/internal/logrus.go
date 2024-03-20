package internal

type Logrus struct {
}

func NewLogrus() *Logrus { return &Logrus{} }

func (l *Logrus) init() {}

func (l *Logrus) Info(s string, args ...interface{}) {
	return
}

func (l *Logrus) Warn(s string, args ...interface{}) {
	return
}

func (l *Logrus) Error(err error, args ...interface{}) {
	return
}

func (l *Logrus) Fatal(err error, args ...interface{}) {
	return
}

func (l *Logrus) Debug(s string, args ...interface{}) {
	return
}

func (l *Logrus) With(args ...interface{}) {
	return
}
