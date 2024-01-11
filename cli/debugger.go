package main

import (
	"github.com/rivo/tview"
)

type (
	Debugger struct {}
)

func NewDebugger() *tview.List {
	debugger := tview.NewList()

	debugger.
		ShowSecondaryText(false).
		SetBorder(true).
		SetTitle("Debugger")

	return debugger
}

func Log(input string) {
	debugger.AddItem(input, "", 0, nil)
}
