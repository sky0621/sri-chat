package chat

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

type client struct {
	ctx    *Ctx
	socket *websocket.Conn
	send   chan *Message // メッセージが送られるチャネル
	room   *room         // このクライアントが参加しているチャットルーム
}

func (c *client) lgr() *logrus.Entry {
	return c.ctx.entry
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		msg := &Message{}
		err := c.socket.ReadJSON(&msg)
		c.lgr().Info(*msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		c.room.message <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	// 1byteずつ処理
	for msg := range c.send {
		// 1byteずつ WebScocket に流し込む
		c.lgr().Info("ブラウザに返すメッセージ：", msg)
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
