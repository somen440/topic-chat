package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/handler"
	"github.com/somen440/topic-chat/src/topic_catalog_service/src/infrastructure"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = 8081
)

var (
	log *logrus.Logger
	srv pb.TopicCatalogServiceServer
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

	srv = handler.NewTopicCatalogServiceServer(
		infrastructure.NewInMemoryTopicRepository(log),
		log,
	)
}

func main() {
	srvPort := os.Getenv("TOPIC_CATALOG_SERVICE_PORT")
	if srvPort == "" {
		srvPort = strconv.Itoa(port)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", srvPort))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterTopicCatalogServiceServer(s, srv)

	log.Infof("starting grpc server at :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
