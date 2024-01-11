package main

import (
	"github.com/rivo/tview"
)

type (
	NavbarItem struct {
		label       string
		description string
		shortcut    rune
		onClick     func()
	}
)

var items = []NavbarItem{
	{
		label:       "VPN",
		description: "",
		shortcut:    0,
		onClick: func() {
			Log("VPN navbar item clicked")
		},
	},
	{
		label:       "Database",
		description: "",
		shortcut:    0,
		onClick: func() {
			Log("Database navbar item clicked")
		},
	},
}

func NewNavbar(title string) *tview.List {
	list := tview.NewList()

	for _, item := range items {
		list.AddItem(item.label, item.description, item.shortcut, item.onClick)
	}

	list.
		ShowSecondaryText(false).
		SetBorder(true).
		SetTitle(title)

	return list
}
