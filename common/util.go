package common

import "syscall/js"

// RemoveElementByID removes an HTML element by ID.
func RemoveElementByID(id string) {
	window := js.Global()
	doc := window.Get("document")
	elem := doc.Call("getElementById", id)
	elem.Call("remove")
}
