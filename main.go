package main

import (
	"log"
	"path/filepath"

	"github.com/leaanthony/mewn"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails"
)

var filename string

func main() {

	homeDirectory, err := homedir.Dir()
	if err != nil {
		log.Fatal("Cannot find home directory! Please file a bug report!")
	}

	filename = filepath.Join(homeDirectory, "todos.json")

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
