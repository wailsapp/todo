package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/leaanthony/mewn"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails"
)

var filename string

func loadTodos() (string, error) {
	result := "[]"
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return "[]", err
		}
		result = string(data)
	}
	return result, nil
}

func save(todos string) error {
	return ioutil.WriteFile(filename, []byte(todos), 0755)
}

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
	app.Bind(loadTodos)
	app.Bind(save)
	app.Run()
}
