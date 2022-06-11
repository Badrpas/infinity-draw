package main

import (
	. "github.com/badrpas/infinity-draw/server/pkg/handlers"
	. "github.com/badrpas/infinity-draw/server/pkg/server"
	. "github.com/badrpas/infinity-draw/server/pkg/wsclient"
)

func main() {
	StartServer(ServerConfig{
		BindAddress: "0.0.0.0:3300",
		RootHandler: RootHandler,
		WsHandler:   GetWebsocketHandler(WsClientFactory),
	})
}
