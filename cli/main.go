// https://gist.github.com/rivo/2893c6740a6c651f685b9766d1898084
// https://github.com/cosmtrek/air
package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

var (
	app      *tview.Application
	pages    *tview.Pages
	debugger *tview.List
)

func main() {
	app = tview.NewApplication()

	setup(app)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func setup(app *tview.Application) {
	placeholder := tview.NewList().ShowSecondaryText(false)

	databases := tview.NewList().ShowSecondaryText(false)
	databases.
		SetBorder(true).
		SetTitle("Databases")

	// Populate dummy data
	for i := 1; i <= 100; i++ {
		databases.AddItem(fmt.Sprintf("asgard-%d", i), "", 0, func(idx int) func() {
			return func() {
				// fmt.Println(fmt.Sprintf("asgard-%d selected", idx))
			}
		}(i))
	}

	navbar := NewNavbar("Menu")

	body := tview.NewFlex().
		AddItem(navbar, 0, 10, true).
		AddItem(databases, 0, 40, false).
		AddItem(placeholder, 0, 50, false)

	debugger = NewDebugger() 

	// Create the layout
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(body, 0, 3, true).
		AddItem(debugger, 0, 1, false)

	pages := tview.NewPages().AddPage("shell", layout, true, true)
	app.SetRoot(pages, true)
}
