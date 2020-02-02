package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/somen440/topic-chat/src/auth_service/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = 8081
)

var (
	users []*pb.User
	log   *logrus.Logger
)

func init() {
	log = logrus.New()
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
	srvPort := os.Getenv("AUTH_SERVICE_PORT")
	if srvPort == "" {
		srvPort = strconv.Itoa(port)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", srvPort))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &authServiceServer{})

	log.Infof("starting grpc server at :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
