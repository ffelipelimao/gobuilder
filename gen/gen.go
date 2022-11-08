package gen

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

//go:embed "mockbuilder.tmpl"
var MockBuilderTmpl string

type Object struct {
	Name   string
	Fields map[string]string
}

func Start(name string, fields string) error {
	if name == "" {
		return errors.New("missing -n flag value")
	}

	object, err := createObject(name, fields)
	if err != nil {
		return err
	}

	filename := strings.ToLower(fmt.Sprintf("%s.go", object.Name))
	generate(MockBuilderTmpl, filename, object)

	return nil
}

func createObject(name string, fields string) (Object, error) {

	fieldsMapper := make(map[string]string)
	fieldsCloven := strings.Split(fields, ",")

	for _, f := range fieldsCloven {
		fieldsClovenByType := strings.Split(f, "-")
		if len(fieldsClovenByType) < 2 {
			return Object{}, errors.New("there's one field who is not in Name-type")
		}
		key := fieldsClovenByType[0]
		value := fieldsClovenByType[1]
		fieldsMapper[key] = value
	}

	return Object{
		Name:   name,
		Fields: fieldsMapper,
	}, nil
}

func generate(MockBuilderTmpl string, outputFile string, data Object) {
	tmpl := template.Must(template.New("").
		Funcs(sprig.FuncMap()).
		Parse(MockBuilderTmpl))

	var processed bytes.Buffer
	err := tmpl.Execute(&processed, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}

	fmt.Println("Writing file: ", outputFile)
	f, _ := os.Create(outputFile)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}
