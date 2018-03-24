package main

import (
	"github.com/clandry94/plant/pkg/display"
)

func main() {
	window, err := display.NewWindow()
	if err != nil {
		panic(err)
	}
	window.Poll()
}
