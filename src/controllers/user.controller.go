package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"

	"messageapp/models"
	"messageapp/services"
)

type UserController struct {
	svc services.UserService
}

func NewUserController(svc services.UserService) UserController {
	return UserController{svc}
}

//UserPostHandler
func (uc *UserController) UserPostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var u models.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	existingUser, _ := uc.svc.FindUserByName(u.Name)
	if existingUser != nil {
		http.Error(w, "User already exist by that name", 500)
		return
	}

	newUser, err := uc.svc.CreateUser(u)
	if err != nil {
		panic(err)
	}

	jsonString, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}

func (uc *UserController) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var u models.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := mux.Vars(r)
	updatedUser, err := uc.svc.UpdateUser(vars["id"], u)
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonString)
}

func (uc *UserController) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := uc.svc.DeleteUser(vars["id"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode("User Deleted")
}

//UsersGetHandle - Gets all users
func (uc *UserController) UsersGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uc.svc.GetAllUsers()
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))
}

func (uc *UserController) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := uc.svc.FindUserByName(vars["name"])
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))
}

func (uc *UserController) UserFindByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := uc.svc.FindUserById(vars["id"])
	if err != nil {
		panic(err)
	}
	jsonString, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(jsonString))
}
