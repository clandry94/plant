package main

import (
	"flag"
	"fmt"
	"github.com/clandry94/plant/pkg/display"
	"github.com/clandry94/plant/pkg/edit"
)

var versionFlag = flag.Bool("version", false, "Show version of the current build")

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: plant [options] [file]...")
	}

	flag.Parse()

	args := flag.Args()

	redisplay, err := display.NewWindow()
	if err != nil {
		panic("could not make window")
	}
	ctx, err := edit.NewContext(redisplay)
	if err != nil {

		panic("could not load any context")
	}

	if len(args) > 0 {
		fmt.Println(args)
		ctx.Load(args[0])

		for {

		}
	}
}
