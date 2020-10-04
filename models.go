package main

import (
	"fmt"
	achieve_broker_go "gopkg.in/IT108/achieve-broker-go.v0"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type authRouter struct {
	router achieve_broker_go.RouterBase
}

func (receiver *authRouter) RunAction(data *kafka.Message)  {
	fmt.Print("test")
}
