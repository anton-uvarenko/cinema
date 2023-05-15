package main

import (
	"github.com/anton-uvarenko/cinema/user-service/internal/core/database"
	"github.com/anton-uvarenko/cinema/user-service/internal/core/repo"
	"github.com/anton-uvarenko/cinema/user-service/internal/services"
	transportGRPC "github.com/anton-uvarenko/cinema/user-service/internal/transport/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db := database.SetUpConnection()
	r := repo.NewRepo(db)
	s := services.NewServices(r)

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	transportGRPC.SetUpServerControllers(server, s)
	err = server.Serve(listen)
	if err != nil {
		logrus.Fatal(err)
	}
}
