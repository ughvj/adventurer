package main

import (
	"context"
	"log"
	"github.com/ughvj/adventurer/pb"
	"google.golang.org/grpc"
)
func main() {
	log.Println("** Start adventure...")
	conn, err := grpc.Dial("host.docker.internal:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("** Fail to connect the guild...:", err)
	}
	defer conn.Close()
	
	client := pb.NewAuthClient(conn)
	req := &pb.LoginRequest{
		Id: 1,
		Secret: "secret",
	}
	
	res, err := client.Login(context.TODO(), req)
	if err != nil {
		log.Fatal("** Fail to login the guild...:", err)
	}
	log.Println(res.GetToken())
}