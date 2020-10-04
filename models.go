package main

import (
	models "../achieve-models-go"
	"encoding/json"
	achieve_broker_go "gopkg.in/IT108/achieve-broker-go.v0"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type authRouter struct {
	router achieve_broker_go.RouterBase
}

func (receiver *authRouter) RunAction(data *kafka.Message) {
	key := string(data.Key)
	switch key {
	case models.AUTH_REGISTER_KEY:
		req := models.RegisterRequest{}
		json.Unmarshal(data.Value, &req)
		log.Print(req.Username)
	default:
		return
	}


}
