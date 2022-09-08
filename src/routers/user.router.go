package routers

import (
	"github.com/gorilla/mux"
	"messageapp/controllers"
	"messageapp/middleware"
)

type UserRouter struct {
	router *mux.Router
}

func NewUserRouter(controller controllers.UserController) UserRouter {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/users", controller.UsersGetHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controller.UserFindByIdHandler).Methods("GET")
	router.HandleFunc("/user", controller.UserPostHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controller.UserUpdateHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.UserDeleteHandler).Methods("DELETE")
	router.HandleFunc("/users/{name}/name", controller.UserGetHandler).Methods("GET")
	return UserRouter{router}
}

func (r UserRouter) GetRouter() *mux.Router {
	return r.router
}