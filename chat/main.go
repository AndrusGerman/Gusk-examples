package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/AndrusGerman/gusk-server"
)

func main() {
	var handlerWsChat = gusk.NewSocket(wsChat)
	// framework gin
	app := gin.Default()
	app.GET("/chat_ws", gin.WrapF(handlerWsChat))
	app.Static("/chat", "./public")
	app.Run()
}

func wsChat(sk *gusk.Socket) {
	// Se desconecta el usuario
	sk.OnClose = func() {
		sk.Upgrader.SendMasive("sms-del-servidor", "El usuario "+sk.ID+" se desconecto", sk)
	}
	// LLegan los mensajes
	sk.Event("sms-del-cliente", func(dato interface{}) {
		mensaje := fmt.Sprintf("'%s': %v", sk.ID, dato)
		// Enviar a todos los usuarios
		sk.Upgrader.SendMasive("sms-del-servidor", mensaje, sk)
	})
	// Welcome solo al usuario actual
	sk.Send("sms-del-servidor", "Hola tu ID es "+sk.ID)
	// A totos los usuarios no al actual que es sk
	sk.Upgrader.SendMasive("sms-del-servidor", "El usuario "+sk.ID+" se conecto", sk)

	// Finish
	fmt.Println("Disconnect: ", <-sk.Finish)
}
