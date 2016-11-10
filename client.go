package chat

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn
	send   chan *message // メッセージが送られるチャネル
	room   *room         // このクライアントが参加しているチャットルーム
}
