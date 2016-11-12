package chat

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// テンプレート管理用の構造体
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// テンプレートハンドラーをレシーバとするHTTPリクエスト処理関数
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(
		func() {
			t.templ = template.Must(template.ParseFiles(filepath.Join("../../templates", t.filename)))
		},
	)

	data := map[string]interface{}{
		"Host": r.Host,
	}
	// if authCookie, err := r.Cookie("auth"); err == nil {
	// 	data["UserData"] = objx.MustFromBase64(authCookie.Value)
	// }

	// HTTPリクエスト情報を渡したことでHTML上でリクエスト情報が参照可能になる
	t.templ.Execute(w, data) // XXX 本当は戻り値をチェックすべき
}
