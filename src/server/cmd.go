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

type Server struct {
	protobufs.UnimplementedUserServer
	svc services.UserService
}

func (s *Server) DeleteUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	var errorMessage string
	log.Printf("User Id: %d", in.GetUserId())
	existingUser, _ := s.svc.FindUserByUserId(in.GetUserId())
	if existingUser == nil {
		errorMessage = fmt.Sprintf("user by user_id %d does not exist", in.GetUserId())
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	err := s.svc.DeleteUser(existingUser.Id.Hex())
	if err != nil {
		errorMessage = fmt.Sprintf("failed deleting user_id %d reason: %s", in.GetUserId(), err)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	errorMessage = fmt.Sprintf("User %d deleted successfully", in.GetUserId())
	return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
}

func (s *Server) SaveUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	var user models.User
	var errorMessage string
	user.Name = *in.Username
	user.UserId = in.UserId
	if len(user.Name) == 0 {
		errorMessage = "username is blank"
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	log.Printf("Username: %s", user.Name)
	existingUser, _ := s.svc.FindUserByName(user.Name)
	if existingUser != nil {
		errorMessage = fmt.Sprintf("Username %s already exist", user.Name)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	log.Printf("User Id: %d", user.UserId)
	existingUser, _ = s.svc.FindUserByUserId(user.UserId)
	if existingUser != nil {
		errorMessage = fmt.Sprintf("User Id %d already exist", user.UserId)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	newUser, err := s.svc.CreateUser(user)
	if err != nil {
		log.Printf("Failed creating user: %s", err.Error())
		return nil, err
	}
	return &protobufs.UserServiceResponse{Msg: fmt.Sprintf("%s has been created successfully!", newUser.Name)}, nil
}

func (s *Server) UpdateUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	var user models.User
	var errorMessage string
	user.Name = in.GetUsername()
	user.UserId = in.GetUserId()
	if len(user.Name) == 0 {
		errorMessage = "username is blank"
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	log.Printf("Username: %s", user.Name)
	existingUser, _ := s.svc.FindUserByName(user.Name)
	if existingUser != nil {
		errorMessage = fmt.Sprintf("Username %s already exist", user.Name)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	log.Printf("User Id: %d", user.UserId)
	existingUser, _ = s.svc.FindUserByUserId(user.UserId)
	if existingUser == nil {
		errorMessage = fmt.Sprintf("User Id %d does not exist", user.UserId)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	updatedUser, err := s.svc.UpdateUser(existingUser.Id.Hex(), user)
	if err != nil {
		errorMessage = fmt.Sprintf("Failed updating user %s", err)
		log.Println(errorMessage)
		return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
	}
	errorMessage = fmt.Sprintf("%s has been updated successfully", updatedUser.Name)
	return &protobufs.UserServiceResponse{Msg: errorMessage}, nil
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