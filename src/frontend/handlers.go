package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

	data := map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
	}

	if isLoggedIn(r) {
		data["user"] = getLoggedInUser(r)
	}

	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
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

func (fe *frontendServer) JoinHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log = log.WithField("method", r.Method)

	if r.Method == http.MethodGet {
		log.Debug("join")
		if err := templates.ExecuteTemplate(w, "join", map[string]interface{}{
			"session_id": sessionID(r),
			"request_id": r.Context().Value(ctxKeyRequestID{}),
		}); err != nil {
			log.Error(err)
		}
		return
	}

	name := r.FormValue("name")
	if name == "" {
		log.Error("name is empty")
	}
	log.WithField("name", name).Debug("join")

	user, err := pb.NewAuthServiceClient(fe.authSvcConn).
		Join(r.Context(), &pb.JoinRequest{
			Name: name,
		})
	if err != nil {
		log.Error(err)
	}

	if err := loggedIn(w, user); err != nil {
		log.Error(err)
	}
	log.WithField("user", user).Debug("join user")

	redirectIndex(w)
}

func (fe *frontendServer) LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log = log.WithField("method", r.Method)

	if r.Method == http.MethodGet {
		log.Debug("loggedin")
		if err := templates.ExecuteTemplate(w, "login", map[string]interface{}{
			"session_id": sessionID(r),
			"request_id": r.Context().Value(ctxKeyRequestID{}),
		}); err != nil {
			log.Error(err)
		}
		return
	}

	id := r.FormValue("id")
	if id == "" {
		log.Error("id is empty")
	}
	log.WithField("id", id).Debug("logged in")

	user, err := pb.NewAuthServiceClient(fe.authSvcConn).
		LoggedIn(r.Context(), &pb.LoggedInRequest{
			UserId: id,
		})
	if err != nil {
		log.Error(err)
	}
	if err := loggedIn(w, user); err != nil {
		log.Error(err)
	}
	log.WithField("user", user).Debug("logged in user")

	redirectIndex(w)
}

func (fe *frontendServer) SignoutHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id := r.FormValue("id")
	if id == "" {
		log.Error("id is empty")
	}
	log.WithField("id", id).Debug("signout")

	_, err := pb.NewAuthServiceClient(fe.authSvcConn).
		Signout(r.Context(), &pb.SignoutRequest{
			UserId: id,
		})
	if err != nil {
		log.Error(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cookieAuth,
		Value:  "",
		MaxAge: -1,
	})

	redirectIndex(w)
}

func sessionID(r *http.Request) string {
	v := r.Context().Value(ctxKeySessionID{})
	if v != nil {
		return v.(string)
	}
	return ""
}

func isLoggedIn(r *http.Request) bool {
	c, _ := r.Cookie(cookieAuth)
	if c != nil && c.Value != "" {
		return true
	}
	return false
}

func loggedIn(w http.ResponseWriter, user *pb.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:   cookieAuth,
		Value:  base64.StdEncoding.EncodeToString(userJSON),
		MaxAge: cookieMaxAge,
	})
	return nil
}

func getLoggedInUser(r *http.Request) *pb.User {
	if !isLoggedIn(r) {
		panic(fmt.Errorf("not loggedIn"))
	}

	c, _ := r.Cookie(cookieAuth)
	decoded, _ := base64.StdEncoding.DecodeString(c.Value)

	user := &pb.User{}
	json.Unmarshal(decoded, user)

	return user
}

func redirectIndex(w http.ResponseWriter) {
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusFound)
}
