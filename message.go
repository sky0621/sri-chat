package chat

import "time"

// Message ... １つのメッセージを表す。※イベントが「メッセージ」の場合に使う。
type Message struct {
	Name    string
	Message string
	When    time.Time
	// FIXME あとは、ID、アバターのURLくらいか。
}
