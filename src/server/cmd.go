package main

import (
	"fmt"
	"log"
	"messageapp/config"
	"messageapp/database"
	"messageapp/models"
	"messageapp/repositories"
	"messageapp/services"
	"net"

	"github.com/keithyw/pbuf-services/protobufs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{
	protobufs.UnimplementedUserServer
	svc services.UserService
}

func (s *Server) SaveUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserMessage, error) {
	var user models.User
	var errorMessage string
	user.Name = in.Username
	if len(user.Name) == 0 {
		errorMessage = "username is blank"
		log.Println(errorMessage)
		return &protobufs.UserMessage{Username: errorMessage}, nil
	}
	log.Printf("Username: %s", in.Username)
	existingUser, _ := s.svc.FindUserByName(in.Username)
	if existingUser != nil {
		errorMessage = fmt.Sprintf("Username %s already exist", in.Username)
		log.Println(errorMessage)
		return &protobufs.UserMessage{Username: errorMessage}, nil
	}
	newUser, err := s.svc.CreateUser(user)
	if err != nil {
		log.Printf("Failed creating user: %s", err.Error())
		return nil, err
	}
	return &protobufs.UserMessage{Username: fmt.Sprintf("%s has been created successfully!", newUser.Name)}, nil
}

func main() {
	s := Server{}
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed loading config: %s", err)
	}
	mongoClient := database.NewMongoDB(config)
	mongoClient.Init()
	s.svc = services.NewUserService(repositories.NewUserRepository(mongoClient.DB.Collection(config.Collection)))
	fmt.Println("server starting")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
			log.Fatalf("Failed listening: %v", err)
	}

	
	serv := grpc.NewServer()
	reflection.Register(serv)
	protobufs.RegisterUserServer(serv, &s)
	if err := serv.Serve(lis); err != nil {
			log.Fatalf("Failed serving: %s", err)
	}

}