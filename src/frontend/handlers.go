package main

import (
	"html/template"
	"net/http"

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

	if err := templates.ExecuteTemplate(w, "home", map[string]interface{}{}); err != nil {
		log.Error(err)
	}
}
