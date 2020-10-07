package main

import (
	"encoding/json"
	broker "gopkg.in/IT108/achieve-broker-go.v0"
)

func sendAnswer(gateId string, clientId string, ans interface{}) {
	result, _ := json.Marshal(ans)
	broker.WriteMsg(gateId, clientId, string(result))
}
