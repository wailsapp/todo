package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails"
)

type Todos struct {
	Runtime  *wails.Runtime
	Log      *wails.CustomLogger
	Filename string
}

func (t *Todos) Save(todos string) error {
	return ioutil.WriteFile(t.Filename, []byte(todos), 0755)
}

func (t *Todos) SaveAs(todos string) error {
	filename := t.Runtime.Dialog.SelectSaveFile()
	t.Filename = filename
	return t.Save(todos)
	return nil
}

func (t *Todos) LoadTodos() (string, error) {
	result := "[]"
	if _, err := os.Stat(t.Filename); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(t.Filename)
		if err != nil {
			return "[]", err
		}
		result = string(data)
	}
	return result, nil
}

func (t *Todos) LoadTodosFromFile() (string, error) {
	filename := t.Runtime.Dialog.SelectFile()
	oldFilename := t.Filename
	t.Filename = filename
	result, err := t.LoadTodos()
	if err != nil {
		t.Filename = oldFilename
	}
	return result, err
}

func (t *Todos) WailsInit(r *wails.Runtime) error {
	t.Runtime = r
	t.Log = r.Log.New("Todos")
	t.Log.Info("Hi from Todos!")
	homeDirectory, err := homedir.Dir()
	if err != nil {
		t.Log.Error("Cannot find home directory! Please file a bug report!")
		return err
	}

	t.Filename = filepath.Join(homeDirectory, "todos.json")

	return nil
}
