package chat

// Room ... チャットルームを表す。クライアント、ボット、及びイベント受信を管理
type Room struct{}

// NewRoom ...
func NewRoom() *Room {
	return &Room{}
}

// Run ...
func (r *Room) Run() {
	for {
		select {
		// 入室イベント発生
		// 退室イベント発生
		// メッセージ送信イベント発生
		}
	}
}
