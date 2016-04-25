package controllers

import (
	"net"
)

type socketFunctions struct {
}

func (c socketFunctions) Login(conn net.Conn, data interface{}) {
	source := data.(map[string]interface{})
	id := source["id"].(string)

	Socketconnections[id] = conn
}

func (c socketFunctions) SendMessageToClient(data interface{}) {
	source := data.(map[string]interface{})

	to := source["to"].(string)
	toConn := Socketconnections[to]
	message := source["message"].(string)

	toConn.Write([]byte(message))
}
