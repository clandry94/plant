package errors

// general control errors for exchanging from plant to the display

import (
	"fmt"
)

func Exit() error {
	return ExitApplication{"exit"}
}

type ExitApplication struct {
	kind string
}

func (e ExitApplication) Error() string {
	return fmt.Sprintf(e.kind)
}
