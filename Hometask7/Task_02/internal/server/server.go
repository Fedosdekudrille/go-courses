package server

import "net/http"

func MustStartServer() {
	mux := http.NewServeMux()
	fsh := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsh))

	mux.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/index.html")
	}))
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getUserByIdHandler)
	mux.HandleFunc("POST /users/{id}", deleteUserByIdHandler)
	http.ListenAndServe(":5000", mux)
}
