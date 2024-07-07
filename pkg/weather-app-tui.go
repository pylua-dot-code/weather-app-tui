package pkg

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Program() error {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(
			tview.NewBox().
				SetBorder(true).
				SetTitleColor(tcell.ColorWhite).
				SetBorderColor(tcell.ColorWhite).
				SetTitle("|Locale Selector|").
				SetBackgroundColor(tcell.ColorDeepSkyBlue),
			45,
			1,
			false,
		).
		AddItem(
			tview.NewBox().
				SetBorder(true).
				SetTitleColor(tcell.ColorWhite).
				SetBorderColor(tcell.ColorWhite).
				SetTitle("|Weather|").
				SetBackgroundColor(tcell.ColorDeepSkyBlue),
			0,
			1,
			false,
		)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		return (err)
	}
	return (nil)
}
