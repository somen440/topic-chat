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
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
}

func (fe *frontendServer) ListTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

	topics, err := fe.ListTopics(r.Context())
	if err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("topics", topics).Debug("list topics")

	if err := templates.ExecuteTemplate(w, "list_topic", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topics":     topics,
	}); err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
}

func (fe *frontendServer) ViewTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id := mux.Vars(r)["id"]
	if id == "" {
		renderHTTPError(log, r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("id", id).
		Debug("view topic")

	topic, err := fe.GetTopic(r.Context(), id)
	if err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("topic", topic).Debug("view topic")

	if err := templates.ExecuteTemplate(w, "view_topic", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topic":      topic,
	}); err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
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
			renderHTTPError(log, r, w, err, http.StatusInternalServerError)
			return
		}
		return
	}

	name := r.FormValue("name")
	if name == "" {
		renderHTTPError(log, r, w, fmt.Errorf("name is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("name", name).Debug("join")

	user, err := fe.Join(r.Context(), name)
	if err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}

	if err := loggedIn(w, user); err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
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
			renderHTTPError(log, r, w, err, http.StatusInternalServerError)
			return
		}
		return
	}

	userID := r.FormValue("id")
	if userID == "" {
		renderHTTPError(log, r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("userID", userID).Debug("logged in")

	user, err := fe.LoggedIn(r.Context(), userID)
	if err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
	if err := loggedIn(w, user); err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("user", user).Debug("logged in user")

	redirectIndex(w)
}

func (fe *frontendServer) SignoutHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	userID := r.FormValue("id")
	if userID == "" {
		renderHTTPError(log, r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("userID", userID).Debug("signout")

	if err := fe.Signout(r.Context(), userID); err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cookieAuth,
		Value:  "",
		MaxAge: -1,
	})

	redirectIndex(w)
}

func (fe *frontendServer) RoomHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("room")

	if !isLoggedIn(r) {
		log.Debug("not logged in")
		redirectIndex(w)
		return
	}

	topicID := r.URL.Query().Get("topic_id")
	if topicID == "" {
		renderHTTPError(log, r, w, fmt.Errorf("not choice topic"), http.StatusInternalServerError)
		return
	}

	topic, err := fe.GetTopic(r.Context(), topicID)
	if err != nil {
		renderHTTPError(log, r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("topic", topic).Debug("get topic")

	if err := templates.ExecuteTemplate(w, "room", map[string]interface{}{
		"session_id": sessionID(r),
		"request_id": r.Context().Value(ctxKeyRequestID{}),
		"topic":      topic,
	}); err != nil {
		log.Error(err)
	}
}

func (fe *frontendServer) RoomJoinHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("room join")

	if !isLoggedIn(r) {
		log.Debug("not logged in")
		redirectIndex(w)
		return
	}

	user := getLoggedInUser(r)
	topicID := r.FormValue("topic_id")
	if topicID == "" {
		renderHTTPError(log, r, w, fmt.Errorf("topic_id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("topic_id", topicID).
		WithField("user", user).Debug("room join")

	w.Header().Set("location", "/room?topic_id="+topicID)
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) RoomLeftHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("room left")

	w.Header().Set("location", "/topic")
	w.WriteHeader(http.StatusFound)
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

func renderHTTPError(log logrus.FieldLogger, r *http.Request, w http.ResponseWriter, err error, code int) {
	log.WithField("error", err).Error("request error")
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
	templates.ExecuteTemplate(w, "error", map[string]interface{}{
		"session_id":  sessionID(r),
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code)})
}
