package main

import (
	"log"
	"net/http"

	"messageapp/config"
	"messageapp/controllers"
	"messageapp/database"
	"messageapp/repositories"
	"messageapp/routers"
	"messageapp/services"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	mongoClient := database.NewMongoDB(config)
	mongoClient.Init()
	svc := services.NewUserService(repositories.NewUserRepository(mongoClient.DB.Collection(config.Collection)))
	userRouter := routers.NewUserRouter(controllers.NewUserController(svc))
	log.Fatal(http.ListenAndServe(config.Port, userRouter.GetRouter()))
}