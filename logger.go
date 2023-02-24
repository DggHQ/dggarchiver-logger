package logger

import (
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

const (
	InvalidLevel log.Level = iota - 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func init() {
	log.SetHandler(text.New((os.Stderr)))
}

func SetLevel(l log.Level) {
	log.SetLevel(l)
}

func SetHandler(h log.Handler) {
	log.SetHandler(h)
}

func Debugf(str string, v ...interface{}) {
	res := `[%s] ` + str
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000 MST")}, v...)
	log.Debugf(res, v...)
}

func Errorf(str string, v ...interface{}) {
	res := `[%s] ` + str
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000 MST")}, v...)
	log.Errorf(res, v...)
}

func Fatalf(str string, v ...interface{}) {
	res := `[%s] ` + str
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000 MST")}, v...)
	log.Fatalf(res, v...)
}

func Infof(str string, v ...interface{}) {
	res := `[%s] ` + str
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000 MST")}, v...)
	log.Infof(res, v...)
}

func Warnf(str string, v ...interface{}) {
	res := `[%s] ` + str
	v = append([]interface{}{time.Now().Format("2006-01-02 15:04:05.000000 MST")}, v...)
	log.Warnf(res, v...)
}
