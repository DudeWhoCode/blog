package rest

import (
	"log"
	"net/http"
	"time"
)

type middlewareLogger struct {
	handler http.Handler
}

func NewLogger(wrapHandler http.Handler) http.Handler {
	return &middlewareLogger{wrapHandler}
}

func (l *middlewareLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %s", time.Now().UTC().Format(time.StampMilli), r.Method, r.URL.Path)
}
