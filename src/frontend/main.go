package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
)

const (
	addr         = "localhost"
	port         = "8080"
	cookieMaxAge = 60 * 60 * 48

	cookiePrefix    = "topic_"
	cookieSessionID = cookiePrefix + "session-id"
)

type ctxKeySessionID struct{}

// var (
// 	templates = template.Must(template.New("").
// 		Funcs(template.FuncMap{}).
// 		ParseGlob("templates/*.html"),
// )

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	srvPort := os.Getenv("PORT")
	if srvPort == "" {
		srvPort = port
	}
	srvAddr := os.Getenv("LISTEN_ADDR")
	if srvPort == "" {
		srvAddr = addr
	}

	r := mux.NewRouter()
	r.HandleFunc("/_healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	var handler http.Handler = r
	handler = &logHandler{log: log, next: handler}
	handler = ensureSessionID(handler)
	handler = &ochttp.Handler{
		Handler:     handler,
		Propagation: &b3.HTTPFormat{}}

	log.Infof("starting server on " + srvAddr + ":" + srvPort)
	log.Fatal(http.ListenAndServe(srvAddr+":"+srvPort, handler))
}
