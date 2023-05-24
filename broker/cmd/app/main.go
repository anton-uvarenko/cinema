package main

import (
	"github.com/anton-uvarenko/cinema/broker-service/internal/transport/grpc"
	"github.com/anton-uvarenko/cinema/broker-service/internal/transport/http"
	"github.com/sirupsen/logrus"
	"log"
	netHttp "net/http"
)

func main() {
	logrus.Info("connecting to auth")
	clients := grpc.ConnectAllClients()
	httpClient := &netHttp.Client{}

	c := http.NewControllers(clients, httpClient)
	r := http.NewRouter(c)

	handler := r.InitRoutes()

	server := http.InitHttpServer(handler)

	logrus.Info("starting listening on port 80")
	log.Fatal(server.ListenAndServe())
}
