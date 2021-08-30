package views

import (
	"bytes"
	"embed"
	"errors"
	"io/fs"
	"text/template"
)

var (
	//go:embed *
	views embed.FS

	templates = make(map[string]*template.Template)
	libs      = make(map[string][]byte)
)

func init() {
	readDir := func(dir string, callback func(dir string, dirEntry fs.DirEntry)) {
		files, _ := views.ReadDir(dir)
		for _, f := range files {
			callback(dir+"/", f)
		}
	}

	readDir("tmpl", func(dir string, dirEntry fs.DirEntry) {
		t, _ := template.ParseFS(views, dir+dirEntry.Name())
		templates[dirEntry.Name()] = t
	})

	readDir("lib", func(dir string, dirEntry fs.DirEntry) {
		fileBytes, _ := views.ReadFile(dir + dirEntry.Name())
		libs[dirEntry.Name()] = fileBytes
	})
}

func Lib(name string) []byte {
	return libs[name]
}

func Render(templateName string, obj interface{}) ([]byte, error) {
	t, ok := templates[templateName]
	if !ok {
		return nil, errors.New("template not found: " + templateName)
	}
	var buffer bytes.Buffer
	t.Execute(&buffer, obj)
	return buffer.Bytes(), nil
}
