package chat

import "time"

// Message ... １つのメッセージを表す。※イベントが「メッセージ」の場合に使う。
type Message struct {
	Message string
	When    time.Time
	Name    string
	// FIXME あとは、ID、アバターのURLくらいか。
}
