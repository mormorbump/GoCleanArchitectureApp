package main

import (
	"app/infrastructure"
)

func main() {
	// 実行すると、ルーターが http.Server にアタッチされ、
	//HTTP リクエストのリスニングとサーブを開始
	// これはhttp.ListenAndServe(addr, router) のショートカット
	// 注意: このメソッドは、エラーが発生しない限り、呼び出し元の goroutine を無期限にブロックする。
	infrastructure.Router.Run()
}