package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	model "modulo/models"
	"modulo/services"

	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body: " + err.Error()))
		return
	}
	var user model.UserModel
	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body: " + err.Error()))
		return
	}

	user, err = services.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating user: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted, Id: %d", user.ID)))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response, err := services.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Getting all users error: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Encoding users error: " + err.Error()))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decoding id error: " + err.Error()))
		return
	}

	user, err := services.GetUser(uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Getting user error: " + err.Error()))
		return
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Encoding users error: " + err.Error()))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decoding id error: " + err.Error()))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body: " + err.Error()))
		return
	}

	var user model.UserModel
	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid body: " + err.Error()))
		return
	}

	err = services.UpdateUser(uint32(ID), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Update user error: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Decoding id error: " + err.Error()))
		return
	}
	err = services.DeleteUser(uint64(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Delete user error: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
