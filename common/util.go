package common

import "syscall/js"

// RemoveElementByID removes an HTML element by ID.
func removeElementByID(doc js.Value, id string) {
	elem := doc.Call("getElementById", id)
	elem.Call("remove")
}
