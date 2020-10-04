package main

import broker "gopkg.in/IT108/achieve-broker-go.v0"

func main() {
	base := broker.RouterBase{}
	authRouter := &authRouter{base}
	broker.AssignRouter(broker.RouterInterface(authRouter))
	broker.Subscribe([]string{broker.AUTH_TOPIC}, broker.AUTH_TOPIC)
	for {}
}
