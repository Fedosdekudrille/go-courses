package server

import (
	"html/template"
	"net/http"
)

func MustStartServer() {
	err := StartServer()
	if err != nil {
		panic(err)
	}
}

var tmpl *template.Template

func StartServer() error {
	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		return err
	}
	tmpl = t
	mux := http.NewServeMux()
	st := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", st))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/complete/{id}", completeTaskHandler) //bonus, can complete task from web by clicking on ✗

	return http.ListenAndServe(":5000", mux)
}
