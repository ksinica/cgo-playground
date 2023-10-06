package main

import (
	"encoding/json"
	"fmt"
	"os"

	playground "github.com/ksinica/cgo-playground"
)

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<url>")
		return 1
	}

	// "http://ftp.nluug.nl/pub/graphics/blender/demo/movies/Sintel.2010.1080p.mkv"
	info, err := playground.ProbeFormat(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return 1
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "    ")
	e.Encode(info)

	return 0
}

func main() {
	os.Exit(run())
}
