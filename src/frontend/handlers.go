package main

import (
	"html/template"
	"net/http"

	pb "github.com/somen440/topic-chat/src/frontend/pb"

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

func (fe *frontendServer) TopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Info("topic")

	resp, err := pb.NewTopicCatalogServiceClient(fe.topicCatalogSvcConn).
		ListTopics(r.Context(), &pb.Empty{})
	if err != nil {
		log.Error(err)
	}

	topics := resp.GetTopics()
	log.WithField("topics", topics).Debug("list topics")

	if err := templates.ExecuteTemplate(w, "topic", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topics":     topics,
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
