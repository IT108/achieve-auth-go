package main

import (
	"encoding/json"
	broker "gopkg.in/IT108/achieve-broker-go.v0"
	models "gopkg.in/IT108/achieve-models-go.v0"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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
		register(req)
		break
	case models.AUTH_AUTHENTICATE_KEY:
		req := models.AuthenticateRequest{}
		json.Unmarshal(data.Value, &req)
		authenticate(req)
		break
	case models.AUTH_AUTHORIZE_KEY:
		req := models.AuthorizeRequest{}
		json.Unmarshal(data.Value, &req)
		authorize(req)
		break
	case models.AUTH_ISREGISTERED_KEY:
		req := models.IsRegisteredRequest{}
		json.Unmarshal(data.Value, &req)
		isRegistered(req)
		break
	default:
		return
	}
}
