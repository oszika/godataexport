package main

//go:generate dataexport -i data -o data.go -m $MODE

import (
	"log"
	"os"
	"text/template"
)

func main() {
	value, ok := Get("data/dir/value")
	if !ok {
		log.Fatal("Can't get data/dir/value")
	}

	content, ok := Get("data/template.tmpl")
	if !ok {
		log.Fatal("Can't get data/template.tmpl")
	}

	t := template.Must(template.New("").Parse(string(content)))

	if err := t.Execute(os.Stdout, string(value)); err != nil {
		log.Fatal(err)
	}
}
