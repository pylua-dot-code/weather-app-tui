package pkg

import (
	"github.com/rivo/tview"
)

func Program() error {
	app := tview.NewApplication()
	flex := tview.NewFlex()
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		return (err)
	}
	return (nil)
}
