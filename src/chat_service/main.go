package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	pb "github.com/somen440/topic-chat/src/chat_service/pb"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

var (
	port = 8083
	log  *logrus.Logger
)

func init() {
	log = logrus.New()
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
}

func main() {
	ctx := context.Background()

	srvPort := os.Getenv("CHAT_SERVICE_PORT")
	if srvPort == "" {
		srvPort = strconv.Itoa(port)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", srvPort))
	if err != nil {
		log.Fatal(err)
	}

	srv := newChatServiceServer()
	srv.topicCatalogSvcAddr = os.Getenv("TOPIC_CATALOG_SERVICE_ADDR")
	srv.authSvcAddr = os.Getenv("AUTH_SERVICE_ADDR")

	mustConnGRPC(ctx, &srv.topicCatalogSvcConn, srv.topicCatalogSvcAddr)
	mustConnGRPC(ctx, &srv.authSvcConn, srv.authSvcAddr)

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, srv)

	log.Infof("starting grpc server at :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(err)
	}
}
