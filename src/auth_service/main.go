package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"
	"github.com/somen440/topic-chat/src/auth_service/src/domain"
	"github.com/somen440/topic-chat/src/auth_service/src/handler"
	"github.com/somen440/topic-chat/src/auth_service/src/infrastructure"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = 8081
)

var (
	srv pb.AuthServiceServer
	log *logrus.Logger
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

	srv = handler.NewAuthServiceServer(
		infrastructure.NewInMemoryUserRepository(
			domain.CreateMockUserMap(),
			log,
		), log,
	)
}

func main() {
	srvPort := os.Getenv("AUTH_SERVICE_PORT")
	if srvPort == "" {
		srvPort = strconv.Itoa(port)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", srvPort))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, srv)

	log.Infof("starting grpc server at :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
