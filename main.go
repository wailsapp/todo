package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Todos",
		Resizable: true,
		JS:        js,
		CSS:       css,
		Colour:    "rgb(245, 245, 245)",
	})

	todoStruct := &Todos{}

	app.Bind(todoStruct)
	app.Run()
}
