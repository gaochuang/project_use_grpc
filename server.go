package main

import (
	"context"
	"fmt"
	"net"
	"user"

	"google.golang.org/grpc"
)

type UserServiceStruct struct {
}

type UserServiceInterface interface {
	GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error)
}

func NewUserService() UserServiceInterface {
	return &UserServiceStruct{}
}

func (userService *UserServiceStruct) GetUser(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	response := &user.UserResponse{
		Id:   req.Id,
		Name: "hellp world",
	}
	return response, nil
}

func main() {
	server, err := net.Listen("tcp", ":8899")
	if err != nil {
		panic(err)
	}

	fmt.Println("listen on local:8899")

	var userService UserServiceInterface

	grpcServer := grpc.NewServer()

	userService = NewUserService()

	user.RegisterUserServiceServer(grpcServer, userService)

	err = grpcServer.Serve(server)
	if err != nil {
		panic(err)
	}

}
