package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// room ... チャットルームを表す。クライアント、ボット、及びイベント受信を管理
type room struct {
	join    chan *client     // 入室しようとしているクライアント
	leave   chan *client     // 退室しようとしているクライアント
	clients map[*client]bool // 在室状態の全クライアント

}

// newRoom ...
func newRoom() *room {
	return &room{}
}

// run ...
func (r *room) run() {
	for {
		select {
		// 入室イベント発生
		// 退室イベント発生
		// メッセージ送信イベント発生
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// WebSocketを使うには、HTTP接続をアップグレードする必要がある
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// chat.html内のJSコードによる「/room」へのHTTPリクエストが来る（新しいクライアントが訪れる）たびに呼ばれるメソッド
// room の参照(*room)を http.Handler 型に適合させる。（※同じシグネチャを持つ ServeHTTP メソッドを追加するだけ）
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil) // HTTP接続をアップグレードしてソケット生成
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	// 構造体「クライアント」を初期化
	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
	}

	r.join <- client // 生成したクライアントをチャットルームの入室用チャネル（join）に投入！

	defer func() { r.leave <- client }() // クライアントの終了時に退室の処理を行う。ユーザがいなくなった際のクリーンナップ。

	// go client.write() // ゴルーチン実行

	// client.read() // メインスレッド上で実行。接続が保持されたまま、終了指示が出るまで他の処理をブロックする。
}
