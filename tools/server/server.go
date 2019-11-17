package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

// Args are in order
// go run server <directory> <port> <open | bool | optional>
func main() {

	dir := os.Args[1]
	port := os.Args[2]
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", http.StripPrefix("/", fs))
	log.Printf("Serving %v on port %v...\n", dir, port)
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func() {
		http.ListenAndServe(":"+port, nil)
		waitGroup.Done()
	}()

	url := "http://localhost:" + port
	log.Printf("Started at url %v\n", url)

	if len(os.Args) > 3 && os.Args[3] == "true" {
		openBrowser(url)
	}

	waitGroup.Wait()
}

func openBrowser(url string) {
	var err error
	log.Printf("Opening url %v\n", url)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
