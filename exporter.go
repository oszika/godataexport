package main

import (
	"os"
	"text/template"
)

type Handler interface {
	AddFile(file *os.File) error
	GetTemplate() string
	GetData() interface{}
}

type Exporter struct {
	exporter Handler
}

func (e *Exporter) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	if stat.IsDir() {
		names, err := file.Readdirnames(-1)
		if err != nil {
			return err
		}

		for _, name := range names {
			e.Read(path + "/" + name)
		}
	} else {
		e.exporter.AddFile(file)
	}

	return nil
}

func (e *Exporter) Write(output string) error {
	w, err := os.Create(output)
	if err != nil {
		return err
	}

	defer w.Close()

	return template.Must(template.New("").Parse(e.exporter.GetTemplate())).Execute(w, e.exporter.GetData())
}
