package controllers

import (
	"html/template"
	"net/http"

	"github.com/luanlouzada/golangweb/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What is the meaning of life?",
			Answer:   "42",
		}, {
			Question: "What is the meaning of life?",
			Answer:   "42",
		}, {
			Question: "What is the meaning of life?",
			Answer:   "42",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
