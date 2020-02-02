package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

var (
	port   = 8081
	topics []*pb.Topic
	log    *logrus.Logger
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

	parseTopics()
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
	pb.RegisterTopicCatalogServiceServer(s, &topicCatalogServiceServer{})

	log.Infof("starting grpc server at :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parseTopics() {
	fileName := "./topic_catalog.yaml"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &topics)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
