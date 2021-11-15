package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

//Add CORS header middleware to the response
func corsMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func (e *Env) loggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...

		e.Logger.Log.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"method":      r.Method,
			"url":         r.URL.Path,
		}).Info("Request")

		next.ServeHTTP(w, r)
	})
}

func (e *Env) profilerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// time the duration of the request
		start := time.Now()
		defer func() {
			e.Logger.Log.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"method":      r.Method,
				"url":         r.URL.Path,
				"duration":    time.Since(start).Microseconds(),
			}).Trace("Request duration")
		}()

		next.ServeHTTP(w, r)
	})
}
