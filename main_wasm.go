//go:build js

package main

import (
	"syscall/js"
)

func main() {
	run(js.Global().Get("wasm_demo").String())
}
