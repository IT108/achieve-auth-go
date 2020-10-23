package main

import (
	"encoding/json"
	methods "github.com/IT108/achieve-auth-go/methods"
	broker "github.com/IT108/achieve-broker-go"
	"github.com/IT108/achieve-models-go/auth"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type authRouter struct {
	router broker.RouterBase
}

func (receiver *authRouter) RunAction(data *kafka.Message) {
	key := string(data.Key)
	switch key {
	case auth.AUTH_REGISTER_KEY:
		req := auth.RegisterRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.Register(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_AUTHENTICATE_KEY:
		req := auth.AuthenticateRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.Authenticate(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_AUTHORIZE_KEY:
		req := auth.AuthorizeRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.Authorize(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_IS_EMAIL_REGISTERED_KEY:
		req := auth.IsEmailRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsEmailRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_IS_USER_REGISTERED_KEY:
		req := auth.IsUserRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsUserRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_ISREGISTERED_KEY:
		req := auth.IsRegisteredRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.IsRegistered(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	case auth.AUTH_SIGNIN_KEY:
		req := auth.SigninRequest{}

		json.Unmarshal(data.Value, &req)
		answer := methods.SignIn(req)

		sendAnswer(answer.GateId, answer.Sender, &answer)
		break
	default:
		return
	}
}
