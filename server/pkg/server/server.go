package server

import (
	"log"
	. "net/http"
)

type ServerConfig struct {
	WsHandler   func(ResponseWriter, *Request)
	RootHandler func(ResponseWriter, *Request)

	BindAddress string
}

func StartServer(c ServerConfig) {
	HandleFunc("/ws", c.WsHandler)
	HandleFunc("/", c.RootHandler)

	log.Fatal(ListenAndServe(c.BindAddress, nil))
}
