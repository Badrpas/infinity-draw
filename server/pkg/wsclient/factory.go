package wsclient

import "github.com/badrpas/infinity-draw/server/pkg/handlers"

func WsClientFactory() handlers.WsApi {
	client := &WsClient{}

	return client
}
