package chat

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

// room ... チャットルームを表す。クライアント、ボット、及びイベント受信を管理
// FIXME bot 分も追加。
// FIXME join/leave/message を１つのイベントにしてタイプによって判別させる？
type room struct {
	ctx     *Ctx
	join    chan *client     // 入室しようとしているクライアント
	leave   chan *client     // 退室しようとしているクライアント
	clients map[*client]bool // 在室状態の全クライアント
	message chan *Message    // クライアントのメッセージ
}

func (r *room) lgr() *logrus.Entry {
	return r.ctx.entry
}

// newRoom ...
func newRoom(ctx *Ctx) *room {
	return &room{
		ctx:     ctx,
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		message: make(chan *Message),
	}
}

// FIXME config.toml に逃がす。
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// WebSocketを使うには、HTTP接続をアップグレードする必要がある
var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

// chat.html内のJSコードによる「/room」へのHTTPリクエストが来る（新しいクライアントが訪れる）たびに呼ばれるメソッド
// room の参照(*room)を http.Handler 型に適合させる。（※同じシグネチャを持つ ServeHTTP メソッドを追加するだけ）
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.lgr().Info("クライアントからリクエストあり！")
	socket, err := upgrader.Upgrade(w, req, nil) // HTTP接続をアップグレードしてソケット生成
	if err != nil {
		r.lgr().Fatal("ServeHTTP:", err)
		return
	}

	// 構造体「クライアント」を初期化
	client := &client{
		ctx:    r.ctx,
		socket: socket,
		send:   make(chan *Message, messageBufferSize),
		room:   r,
	}

	r.join <- client // 生成したクライアントをチャットルームの入室用チャネル（join）に投入！

	defer func() { r.leave <- client }() // クライアントの終了時に退室の処理を行う。ユーザがいなくなった際のクリーンナップ。

	go client.write() // ゴルーチン実行

	client.read() // メインスレッド上で実行。接続が保持されたまま、終了指示が出るまで他の処理をブロックする。
}

// run ...
func (r *room) run() {
	r.lgr().Info("チャットルーム起動！")
	for {
		// FIXME クライアントとボットとで分ける！
		select {
		case client := <-r.join: // 入室イベント発生！
			r.clients[client] = true
			r.lgr().Info("クライアントが入室しました。")
			// 退室イベント発生
		case client := <-r.leave: // 退室イベント発生！
			delete(r.clients, client) // 在室状態から消す
			close(client.send)        // 消したクライアントの送信チャネルを閉じる
			r.lgr().Info("クライアントが退室しました。")
		case msg := <-r.message: // メッセージ送信イベント発生
			r.lgr().Info("メッセージを受信しました。：", msg.Message)
			for client := range r.clients {
				select {
				case client.send <- msg: // １人１人のクライアントのチャネルにメッセージを流し込む
					r.lgr().Infof(" -- クライアントに送信されました")
				default: // 【転送失敗】
					delete(r.clients, client) // 在室状態から消す
					close(client.send)        // 消したクライアントの送信用チャネルを閉じる
					r.lgr().Infof(" -- 送信に失敗しました。クライアントをクリーンナップします。")
				}
			}
		}
	}
}
