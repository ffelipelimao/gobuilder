package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
)

type Object struct {
	Name   string
	Fields map[string]string
}

func main() {

	object := Object{
		Name: "Customer",
		Fields: map[string]string{
			"Name":    "string",
			"Phone":   "int64",
			"Address": "string",
		},
	}

	processTemplate("mockbuilder.tmpl", "builder.go", object)
}

func processTemplate(fileName string, outputFile string, data Object) {
	tmpl := template.Must(template.New("gobuilder").Funcs(sprig.FuncMap()).ParseFiles(fileName))

	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}

	outputPath := "./mockbuilder/" + outputFile
	fmt.Println("Writing file: ", outputPath)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}
