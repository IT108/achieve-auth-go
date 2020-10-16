package main

import (
	"encoding/json"
	broker "github.com/IT108/achieve-broker-go"
	models "github.com/IT108/achieve-models-go"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type authRouter struct {
	router broker.RouterBase
}

func (receiver *authRouter) RunAction(data *kafka.Message) {
	key := string(data.Key)
	switch key {
	case models.AUTH_REGISTER_KEY:
		req := models.RegisterRequest{}
		json.Unmarshal(data.Value, &req)
		answer := register(req)
		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_AUTHENTICATE_KEY:
		req := models.AuthenticateRequest{}
		json.Unmarshal(data.Value, &req)
		answer := authenticate(req)
		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_AUTHORIZE_KEY:
		req := models.AuthorizeRequest{}
		json.Unmarshal(data.Value, &req)
		answer := authorize(req)
		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_ISREGISTERED_KEY:
		req := models.IsRegisteredRequest{}
		log.Println(req.GateId)
		json.Unmarshal(data.Value, &req)
		answer := isRegistered(req)
		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	default:
		return
	}
}
