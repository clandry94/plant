package main

import (
	"flag"
	"fmt"
	"github.com/clandry94/plant/pkg/cmd"
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

	window, err := display.NewWindow()
	if err != nil {
		panic("could not make window")
	}

	ctx, err := edit.NewContext()
	if err != nil {
		panic("could not load any context")
	}

	if len(args) > 0 {
		fmt.Println(args)
		ctx.Load(args[0])

		window.Screen().ShowCursor(0, 0)
		for {
			ev := window.Screen().PollEvent()
			cmd.Handle(ev, ctx.CurrentBuffer())
			window.SetCursor(ctx.CurrentBuffer().GetCursor().Col(), ctx.CurrentBuffer().GetCursor().Line())
			window.Refresh(ctx.CurrentBuffer().GetContents())
			window.RefreshStatusLine(ctx.CurrentBuffer().Status())
			//window.Redisplay()
		}
	}
}
