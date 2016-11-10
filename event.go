package chat

const (
	// evtJoin ... 入室イベント
	evtJoin = iota
	// evtLeave ... 退室イベント
	evtLeave
	// evtMessage ... メッセージ送信イベント
	evtMessage
)

// event ...
type event struct {
	evt int
}
