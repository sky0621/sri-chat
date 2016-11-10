package chat

import "time"

// message ... １つのメッセージを表す。※イベントが「メッセージ」の場合に使う。
type message struct {
	msg  string
	when time.Time
	// FIXME あとは、ID、名前、アバターのURLくらいか。
}
