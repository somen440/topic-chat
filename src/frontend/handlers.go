package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	pb "github.com/somen440/topic-chat/src/common/pb"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	templates = template.Must(template.New("").
		Funcs(template.FuncMap{}).
		ParseGlob("templates/*.html"))
)

type emptyData map[string]interface{}

var empty = emptyData{}

func (fe *frontendServer) HomeHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Info("home")
	fe.execute(w, r, "home", empty)
}

func (fe *frontendServer) ListTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

	topics, err := fe.ListTopics(r.Context())
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("topics", topics).Debug("list topics")

	fe.execute(w, r, "list_topic", map[string]interface{}{"topics": topics})
}

func (fe *frontendServer) ViewTopicHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		renderHTTPError(r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	if id == 0 {
		renderHTTPError(r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("id", id).
		Debug("view topic")

	topic, err := fe.GetTopic(r.Context(), id)
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("topic", topic).Debug("view topic")

	fe.execute(w, r, "view_topic", map[string]interface{}{
		"topic": topic,
	})
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
			renderHTTPError(r, w, err, http.StatusInternalServerError)
			return
		}
		return
	}

	name := r.FormValue("name")
	if name == "" {
		renderHTTPError(r, w, fmt.Errorf("name is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("name", name).Debug("join")

	user, err := fe.Join(r.Context(), name)
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}

	if err := loggedIn(w, user); err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
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
		fe.execute(w, r, "login", empty)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	if userID == 0 {
		renderHTTPError(r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("userID", userID).Debug("logged in")

	user, err := fe.LoggedIn(r.Context(), userID)
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	if err := loggedIn(w, user); err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	log.WithField("user", user).Debug("logged in user")

	redirectIndex(w)
}

func (fe *frontendServer) SignoutHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	userID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
	if userID == 0 {
		renderHTTPError(r, w, fmt.Errorf("id is empty"), http.StatusInternalServerError)
		return
	}
	log.WithField("userID", userID).Debug("signout")

	if err := fe.Signout(r.Context(), userID); err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cookieAuth,
		Value:  "",
		MaxAge: -1,
	})

	redirectIndex(w)
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
		renderHTTPError(r, w, fmt.Errorf("topic_id is empty"), http.StatusInternalServerError)
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

func (fe *frontendServer) execute(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if isLoggedIn(r) {
		data["user"] = getLoggedInUser(r)
	}

	data["session_id"] = sessionID(r)
	data["request_id"] = r.Context().Value(ctxKeyRequestID{})

	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		renderHTTPError(r, w, err, http.StatusInternalServerError)
		return
	}
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

func renderHTTPError(r *http.Request, w http.ResponseWriter, err error, code int) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.WithField("error", err).Error("request error")
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
	templates.ExecuteTemplate(w, "error", map[string]interface{}{
		"session_id":  sessionID(r),
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code)})
}
