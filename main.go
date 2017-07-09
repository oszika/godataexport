package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var input = flag.String("i", "", "input <file | directory>")
	var output = flag.String("o", "", "output <go source file | destination export path>")
	var mode = flag.String("m", "", "mode <source | path>")
	var packageName = flag.String("p", "main", "package name")

	flag.Parse()

	if *input == "" || *output == "" {
		flag.PrintDefaults()
		log.Fatal("Missing i/o files")
	}

	var handler Handler

	if strings.EqualFold(*mode, "source") {
		handler = NewExporter2Src(*packageName)
	} else if strings.EqualFold(*mode, "path") {
		handler = NewExporter2Path(*packageName)
	} else {
		flag.PrintDefaults()
		log.Fatal("Invalid mode")
	}

	e := &Exporter{handler}

	if err := e.Read(*input); err != nil {
		log.Fatal(err)
	}

	if err := e.Write(*output); err != nil {
		log.Fatal(err)
	}
}
