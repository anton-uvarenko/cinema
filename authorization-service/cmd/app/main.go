package main

import (
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/database"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/core/repo"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/services"
	"github.com/anton-uvarenko/cinema/authorization-service/internal/transport/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db := database.SetUpConnection()
	r := repo.NewRepo(db)

	s := services.NewService(r)

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	err = pkg.InitGrpcConnection()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	rpc.SetUpServerControllers(server, s)

	logrus.Info("starting server ...")
	err = server.Serve(listen)
	if err != nil {
		logrus.Fatal(err)
	}
}
