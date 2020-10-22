package main

import (
	"encoding/json"
	broker "github.com/IT108/achieve-broker-go"
)

func sendAnswer(gateId string, clientId string, ans interface{}) {
	if gateId == "" || clientId == "" {
		return
	}
	result, _ := json.Marshal(ans)
	broker.WriteMsg(gateId, clientId, string(result))
}
