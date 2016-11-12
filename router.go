package chat

import "net/http"

// Routing ...
func Routing(ctx *Ctx) (exitCode int, err error) {
	r := newRoom(ctx)
	// FIXME goji使用に変える。
	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.run() // チャットルーム開始 -> 入退室やメッセージを待ち受ける

	ctx.entry.Info("Webサーバーを開始します。ポート：", ctx.Port)
	if err := http.ListenAndServe(ctx.Port, nil); err != nil {
		ctx.entry.Fatal("ListenAndServe:", err)
	}
	return ExitCodeOK, nil
}
