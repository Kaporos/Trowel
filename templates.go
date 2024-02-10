package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Template struct {
	name string
}

func basePath() string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to get $HOME, please check if it has a value.")
	}
	complete_path := filepath.Join(path, ".local/share/trowel/")
	if os.MkdirAll(path, 0666) != nil {
		println("Unable to create", complete_path)
		panic("could not recover.")
	}
	return complete_path

}

func listTemplates() []Template {
	entries, err := os.ReadDir(basePath())
	if err != nil {
		println("Unable to retrieve templates")
		return nil
	}
	templates := []Template{}
	for _, e := range entries {
		if e.IsDir() {
			templates = append(templates, Template{
				name: e.Name(),
			})
		}
	}
	return templates
}

func createTemplate(toName string, templateName string) {
	path := filepath.Join(filepath.Join(basePath(), templateName))
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Template", templateName, "not exist.")
		return
	}
	os.Mkdir(toName, 0700)
	copyEntries(entries, path, toName)
}

func registerTemplate(origin string) {
	entries, err := os.ReadDir(origin)
	if err != nil {
		fmt.Println("Unable to read", origin)
		return
	}
	path := filepath.Join(basePath(), filepath.Base(origin))
	os.Mkdir(path, 0700)
	copyEntries(entries, origin, path)
	fmt.Println("Registered", origin)
}

func deleteTemplate(origin string) {
	path := filepath.Join(basePath(), filepath.Base(origin))
	_, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Unable to delete", origin, "are you sure it exists ?")
		return
	}
	os.RemoveAll(path)

}

func copyEntries(entries []fs.DirEntry, path string, destination string) {
	for _, e := range entries {
		name := e.Name()
		ndest := filepath.Join(destination, name)
		if e.IsDir() {
			npath := filepath.Join(path, name)
			new_entries, err := os.ReadDir(npath)
			if err != nil {
				fmt.Println("Unable to read", npath)
				return
			}
			os.Mkdir(ndest, 0700)
			copyEntries(new_entries, npath, ndest)
		} else {
			data, err := os.ReadFile(filepath.Join(path, name))
			if err != nil {
				fmt.Println("Unable to read", name)
				return
			}
			os.WriteFile(ndest, data, 0644)

		}
		fmt.Println("Created", ndest)
	}
}
