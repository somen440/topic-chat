package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/somen440/topic-chat/src/topic_catalog_service/pb"

	"gopkg.in/yaml.v2"
)

var (
	port   = 8081
	topics []*pb.Topic
)

type topicCatalogUsecase struct{}

func (t *topicCatalogUsecase) ListTopics(_ context.Context, _ *pb.Empty) (*pb.ListTopicsResponse, error) {
	return &pb.ListTopicsResponse{
		Topics: topics,
	}, nil
}

func (t *topicCatalogUsecase) GetTopic(_ context.Context, req *pb.GetTopicRequest) (*pb.Topic, error) {
	for _, v := range topics {
		if v.GetId() == req.GetId() {
			return v, nil
		}
	}
	return nil, fmt.Errorf("not found")
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

	fileName := "./topic_catalog.yaml"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &topics)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTopicCatalogServiceServer(s, &topicCatalogUsecase{})

	log.Printf("liten: :%s", srvPort)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
