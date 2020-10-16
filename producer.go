package main

import (
	"encoding/json"
	broker "github.com/IT108/achieve-broker-go"
)

func sendAnswer(gateId string, clientId string, ans interface{}) {
	result, _ := json.Marshal(ans)
	broker.WriteMsg(gateId, clientId, string(result))
}
