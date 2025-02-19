package main

import (
	"syscall/js"
)

func main() {
	// 1. 브라우저 콘솔 출력
	println("WASM 샌드박스 초기화 완료")

	// 2. DOM 요소 제어
	doc := js.Global().Get("document")
	title := doc.Call("createElement", "h1")
	title.Set("innerHTML", "Go WASM 작동 확인!")
	blueArea := doc.Call("getElementById", "wasm-output")
	blueArea.Call("appendChild", title)

	// 3. 자바스크립트 alert 호출
	alert := js.Global().Get("alert")
	alert.Invoke("브라우저 API 정상 호출")

	// 4. 이벤트 루프 유지
	select {}
}
