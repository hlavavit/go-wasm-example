//
package main

import (
	"fmt"
	"log"
	"os"
	"syscall/js"

	"github.com/hlavavit/go-wasm-threejs/three"
)

func main() {
	//print to js console
	fmt.Println("Hello, WebAssembly!")
	// change dom
	doc := js.Global().Get("document")
	app := doc.Call("getElementById", "app")

	app.Set("innerHTML", "Hello, WebAssembly!")

	clonedApp := app.Call("cloneNode", false)
	app.Get("parentNode").Call("replaceChild", clonedApp, app)
	app = clonedApp

	canvas := doc.Call("createElement", "canvas")

	canvas.Set("id", "CursorLayer")
	canvas.Set("width", "800")
	canvas.Set("height", "500")
	canvas.Get("style").Set("zIndex", "8")
	canvas.Get("style").Set("position", "absolute")
	canvas.Get("style").Set("border", "1px solid")

	app.Call("appendChild", canvas)

	_, err := three.Init()
	if err != nil {
		log.Fatalln(err)
	}

	rendererParams := make(map[string]interface{})
	rendererParams["canvas"] = canvas

	renderer := three.NewWebGLRenderer(rendererParams)
	fmt.Printf("renderer = %+v\n", renderer)

	const fov = 75
	const aspect = float32(800) / float32(500) // the canvas default
	const near = 0.1
	const far = 55

	camera := three.NewPerspectiveCamera(fov, aspect, near, far)
	camera.GetPosition().SetZ(2)
	fmt.Printf("camera = %+v\n", camera)

	scene := three.NewScene()
	fmt.Printf("scene = %+v\n", scene)

	geometry := three.NewBoxGeometry(1, 1, 1, 1, 1, 1)
	fmt.Printf("geometry = %+v\n", geometry)

	meshParams := make(map[string]interface{})
	meshParams["color"] = 0x44aa88

	material := three.NewMeshPhongMaterial(meshParams)
	fmt.Printf("material = %+v\n", material)

	cube := three.NewMesh(geometry, material)
	fmt.Printf("cube = %+v\n", cube)

	scene.Add(cube)

	light := three.NewDirectionalLight(0xFFFFFF, 1)
	light.GetPosition().Set(0, 5, 10)
	scene.Add(light)

	exit := make(chan int)
	var render three.RenderFunc = nil
	render = func(time int) {
		seconds := float64(time) * 0.001

		cube.GetRotation().SetX(seconds)
		cube.GetRotation().SetZ(seconds)

		renderer.Render(scene, camera)

		three.RequestAnimationFrame(render)
	}

	three.RequestAnimationFrame(render)

	os.Exit(<-exit)
}
