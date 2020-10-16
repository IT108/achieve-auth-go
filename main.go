package main

import (
	broker "github.com/IT108/achieve-broker-go"
	db "github.com/IT108/achieve-db-go"
	"log"
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

func main() {
	shutdownService()
	db.ConfigureFromEnv()
	db.Init()
	getConfig()
	base := broker.RouterBase{}
	authRouter := &authRouter{base}
	broker.AssignRouter(broker.RouterInterface(authRouter))
	broker.Subscribe([]string{broker.AUTH_TOPIC}, broker.AUTH_TOPIC)

	for {
	}
}
