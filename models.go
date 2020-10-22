package main

import (
	"encoding/json"
	methods "github.com/IT108/achieve-auth-go/methods"
	broker "github.com/IT108/achieve-broker-go"
	models "github.com/IT108/achieve-models-go"
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
		answer := methods.Register(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_AUTHENTICATE_KEY:
		req := models.AuthenticateRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.Authenticate(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_AUTHORIZE_KEY:
		req := models.AuthorizeRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.Authorize(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_IS_EMAIL_REGISTERED_KEY:
		req := models.IsEmailRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsEmailRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_IS_USER_REGISTERED_KEY:
		req := models.IsUserRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsUserRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case models.AUTH_ISREGISTERED_KEY:
		req := models.IsRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break

	default:
		return
	}
}
