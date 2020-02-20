package main

import (
	"github.com/ivansukach/profile-service-example/grpc/protocol"
	"github.com/ivansukach/profile-service-example/grpc/repositories"
	"github.com/ivansukach/profile-service-example/grpc/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	rps, _ := repositories.OpenPostgreSQLRepository()

	log.Info("GRPC-server started")
	s := grpc.NewServer()                         //создали сервер
	srv := &server.Server{}                       //ссылка на пустую структуру
	protocol.RegisterProfileServiceServer(s, srv) //зарегистрировали сервер
	listener, err := net.Listen("tcp", ":1323")   //просто слушаем порт
	s.Serve(listener)
	if err != nil {
		log.Error(err)
	}

	defer rps.CloseDB()
}
