package chat

const (
	// EvtJoin ... 入室イベント
	EvtJoin = iota
	// EvtLeave ... 退室イベント
	EvtLeave
	// EvtMessage ... メッセージ送信イベント
	EvtMessage
)

// Event ...
type Event struct {
	Evt int
}
