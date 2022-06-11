package handlers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

type WsApi interface {
	Handshake(conn *websocket.Conn) error

	DecodeMessage(msgType int, data []byte) error

	OnClose(*websocket.Conn)
}

func GetWebsocketHandler(wsApiFactory func() WsApi) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Upgrade error:", err)
			return
		}

		log.Println("New WS connection")

		api := wsApiFactory()

		defer func(c *websocket.Conn) {
			log.Println("Closing connection")

			api.OnClose(c)
			err := c.Close()
			if err != nil {
				log.Print(err)
			}
		}(c)

		{
			err := api.Handshake(c)
			if err != nil {
				log.Print("Handshake Error:", err)
				return
			}
		}

		log.Println("Handshake passed")

		for {
			msgType, data, err := c.ReadMessage()
			if err != nil {
				log.Print("ReadMessage Error:", err)
				return
			}

			err = api.DecodeMessage(msgType, data)
			if err != nil {
				log.Print("DecodeMessage Error:", err)
				return
			}
		}

	}
}
