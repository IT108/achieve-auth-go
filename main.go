package main

import (
	"github.com/IT108/achieve-auth-go/auth"
	config "github.com/IT108/achieve-auth-go/config"
	"github.com/IT108/achieve-auth-go/methods"
	broker "github.com/IT108/achieve-broker-go"
	db "github.com/IT108/achieve-db-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func shutdownService() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		broker.StopConsumer()
		os.Exit(0)
	}()
}

//func newServer() *auth.AuthServiceServer {
//	s := &methods.Server{}
//	return s
//}

func main() {
	shutdownService()
	db.ConfigureFromEnv()
	broker.ConfigureFromEnv()
	db.Init()
	config.GetConfig()
	//base := broker.RouterBase{}
	//authRouter := &authRouter{base}
	//broker.AssignRouter(broker.RouterInterface(authRouter))
	//broker.Subscribe([]string{broker.AUTH_TOPIC}, broker.AUTH_TOPIC)
	//
	//for {
	//}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := methods.Server{}

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
