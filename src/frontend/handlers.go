package main

import (
	"html/template"
	"net/http"

	pb "github.com/somen440/topic-chat/src/frontend/pb"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	templates = template.Must(template.New("").
		Funcs(template.FuncMap{}).
		ParseGlob("templates/*.html"))
)

func (fe *frontendServer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Info("home")

	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
	}); err != nil {
		log.Error(err)
	}
}

func (fe *frontendServer) ListTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

	resp, err := pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		ListTopics(r.Context(), &pb.Empty{})
	if err != nil {
		log.Error(err)
	}

	topics := resp.GetTopics()
	log.WithField("topics", topics).Debug("list topics")

	if err := templates.ExecuteTemplate(w, "list_topic", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topics":     topics,
	}); err != nil {
		log.Error(err)
	}
}

func (fe *frontendServer) ViewTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id := mux.Vars(r)["id"]
	if id == "" {
		log.Error("id is empty")
	}
	log.WithField("id", id).
		Debug("view topic")

	topic, err := pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		GetTopic(r.Context(), &pb.GetTopicRequest{
			Id: id,
		})
	if err != nil {
		log.Error(err)
	}
	log.WithField("topic", topic).Debug("view topic")

	if err := templates.ExecuteTemplate(w, "view_topic", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topic":      topic,
	}); err != nil {
		log.Error(err)
	}
}

func sessionID(r *http.Request) string {
	v := r.Context().Value(ctxKeySessionID{})
	if v != nil {
		return v.(string)
	}
	return ""
}
