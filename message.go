package chat

import "time"

// Message ... １つのメッセージを表す。※イベントが「メッセージ」の場合に使う。
type Message struct {
	Msg  string
	When time.Time
	// FIXME あとは、ID、名前、アバターのURLくらいか。
}
