package chat

import (
	"log"
	"net/http"
)

// Routing ...
func Routing(ctx *Ctx) (exitCode int, err error) {

	r := newRoom()
	http.Handle("/room", r)

	go r.run() // チャットルーム開始 -> 入退室やメッセージを待ち受ける

	log.Println("Webサーバーを開始します。ポート：", ctx.Port)
	if err := http.ListenAndServe(ctx.Port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	return ExitCodeOK, nil
}
