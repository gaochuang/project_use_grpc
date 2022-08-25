package main

import (
	"context"
	"encoding/json"
	"fmt"
	"user"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8899", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer func() {
		conn.Close()
	}()

	client := user.NewUserServiceClient(conn)
	req := &user.UserRequest{
		Id: 1,
	}

	response, _ := client.GetUser(context.Background(), req)

	resp, err := json.Marshal(response)

	fmt.Printf("%s\n", resp)

}
