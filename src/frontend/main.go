package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"google.golang.org/grpc"
)

const (
	addr         = "localhost"
	port         = "8080"
	cookieMaxAge = 60 * 60 * 48

	cookiePrefix    = "topicchat_"
	cookieSessionID = cookiePrefix + "session-id"
	cookieAuth      = cookiePrefix + "auth"
)

type ctxKeySessionID struct{}

type frontendServer struct {
	topicCatalogSvcAddr string
	topicCatalogSvcConn *grpc.ClientConn

	authSvcAddr string
	authSvcConn *grpc.ClientConn
}

func main() {
	ctx := context.Background()
	log := logrus.New()
	if os.Getenv("DEBUG") == "on" {
		log.Level = logrus.DebugLevel
	}
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	srvPort := os.Getenv("FRONTEND_PORT")
	if srvPort == "" {
		srvPort = port
	}
	srvAddr := os.Getenv("FRONTEND_LISTEN_ADDR")
	if srvPort == "" {
		srvAddr = addr
	}

	srv := new(frontendServer)
	srv.topicCatalogSvcAddr = os.Getenv("TOPIC_CATALOG_SERVICE_ADDR")
	srv.authSvcAddr = os.Getenv("AUTH_SERVICE_ADDR")

	mustConnGRPC(ctx, &srv.topicCatalogSvcConn, srv.topicCatalogSvcAddr)
	mustConnGRPC(ctx, &srv.authSvcConn, srv.authSvcAddr)

	r := mux.NewRouter()
	r.HandleFunc("/", srv.HomeHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/topic", srv.ListTopicHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/topic/{id}", srv.ViewTopicHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/join", srv.JoinHandler).Methods(http.MethodGet, http.MethodHead, http.MethodPost)
	r.HandleFunc("/login", srv.LoggedInHandler).Methods(http.MethodGet, http.MethodHead, http.MethodPost)
	r.HandleFunc("/signout", srv.SignoutHandler).Methods(http.MethodHead, http.MethodPost)
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

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
