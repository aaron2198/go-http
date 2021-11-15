package logger

import (
	"errors"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type VTSBLogger struct {
	Log *logrus.Logger
}

// CreateLogger create a new custom logger at the given location and level (default: stdout, info)
func CreateLogger(loc string, level logrus.Level) *VTSBLogger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logger.SetLevel(level)
	logger.Formatter = &logrus.JSONFormatter{}
	if len(loc) > 0 {
		f, err := os.OpenFile(loc, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		logger.Out = f
	} else {
		logger.Out = os.Stdout
	}
	return &VTSBLogger{logger}
}

// RequestError is a default error handler for requests, parses http.Request into logrus fields for accurate analysis and a standard interface
func (l *VTSBLogger) RequestError(r *http.Request, err error) {
	l.Log.WithFields(logrus.Fields{
		"request": r.URL.Path,
		"method":  r.Method,
		"err":     err.Error(),
	}).Error("Request error")
}

// RequestDefaultError wraps the RequestError with an error message to the writer, if error is nil, sends method not allowed.
func (l *VTSBLogger) RequestDefaultError(r *http.Request, w *http.ResponseWriter, err error) {
	if err == nil {
		err = errors.New("method not allowed")
	}
	l.RequestError(r, err)
	http.Error(*w, err.Error(), http.StatusInternalServerError)
}
