package controllers

import (
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/dennypradipta/cariliquid/models"
	"github.com/dennypradipta/cariliquid/responses"
)

// GetUsers returns all users
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get User Model
	user := &models.User{}

	fmt.Println(time.Now())

	// Find User in DB
	users, err := user.FindUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUserByID returns single user
func (server *Server) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Read query params
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Get User Model
	user := &models.User{}

	// Find User in DB by ID
	user, err := user.FindUserByID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// CreateUser returns created user
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Get User Model
	user := models.User{}

	// Parse JSON then input the value to User
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Prepare new User but Validate the fields first
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create the user
	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}

// UpdateUserByID returns created user
func (server *Server) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	// Read query params
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Get User Model
	user := models.User{}

	// Parse JSON then input the value to User
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Validate the fields first
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create the user
	userUpdated, err := user.UpdateUserByID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, userUpdated.ID))
	responses.JSON(w, http.StatusCreated, userUpdated)
}

// DeleteUserByID returns number of rows affected
func (server *Server) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Read query params
	vars := mux.Vars(r)
	userID := vars["userID"]

	// Get User Model~
	user := models.User{}

	// Delete the user
	_, err := user.DeleteUserByID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.RequestURI, userID))
	responses.JSON(w, http.StatusNoContent, nil)
}