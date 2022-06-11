package wsclient

import (
	"errors"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	conn *websocket.Conn
}

func (client *WsClient) Handshake(conn *websocket.Conn) error {
	msgType, data, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	if msgType == websocket.TextMessage && data != nil {
		message := string(data)
		if message == "hello" {
			client.conn = conn
			return nil
		}
	}

	return errors.New("Handshake failed")
}

func (client *WsClient) SendMessage(msgType int, data []byte) error {
	if client.conn == nil {
		return errors.New("WsClient is not initialized with handshake")
	}

	return client.conn.WriteMessage(msgType, data)
}

func (client *WsClient) SendBytes(message []byte) error {
	return client.SendMessage(websocket.BinaryMessage, message)
}
func (client *WsClient) SendString(message string) error {
	return client.SendMessage(websocket.TextMessage, []byte(message))
}

func (client *WsClient) OnClose(conn *websocket.Conn) {
}
