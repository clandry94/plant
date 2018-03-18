package display_test

import(
	"github.com/clandry94/plant/display"
	"testing"
)

func TestNewWindow(t *testing.T) {
	_, err := display.NewWindow()
	if err != nil {
		t.Error(err)
	}

}
