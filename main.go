//
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	//print to js console
	fmt.Println("Hello, WebAssembly!")
	// change dom
	doc := js.Global().Get("document")
	body := doc.Call("getElementById", "app")
	body.Set("innerHTML", "Hello, WebAssembly!")
}
