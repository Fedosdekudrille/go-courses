package server

import (
	"Task_02/pkg/udb"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := udb.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	fmt.Println(sid)
	id, err := strconv.Atoi(sid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	user, err := udb.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func deleteUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	sid := r.PathValue("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = udb.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = udb.DeleteUser(id)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	users, err := udb.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if json.NewEncoder(w).Encode(users) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
