package main

import (
	"CarLibrary/testdrive/config"
	"CarLibrary/testdrive/model"
	"CarLibrary/testdrive/testdrive"
	"fmt"
	pb "github.com/CarLibrary/proto/testdrive"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	//加载配置
	config.Config()
	model.InitMYSQL()
	fmt.Println("ok")

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestDriveServerServer(s, &testdrive.TestDriveServerServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
