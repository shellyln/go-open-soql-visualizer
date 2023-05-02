//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	"github.com/shellyln/go-open-soql-parser/soql/parser"
	"github.com/shellyln/go-open-soql-visualizer/soql/visualizer"
)

func visualizeSoql(this js.Value, args []js.Value) interface{} {
	src := ""
	if 0 < len(args) {
		src = args[0].String()
	}

	parsedQuery, err := parser.Parse(src)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	s := visualizer.Visualize(parsedQuery)

	return js.ValueOf(s)
}

func getVersion(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(Version)
}

func main() {
	println("Go WebAssembly Initialized")

	js.Global().Set("visualizeSoql", js.FuncOf(visualizeSoql))
	js.Global().Set("getVersion", js.FuncOf(getVersion))

	select {}
}
