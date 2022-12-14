package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ffelipelimao/gobuilder/cmd/gen/tmpls"
)

type Object struct {
	Name   string
	Fields map[string]string
}

func Start(name string, fields string, destination string) {
	if name == "" {
		log.Fatalf("Missing -n flag value")
	}

	object := createObject(name, fields)

	filename := strings.ToLower(fmt.Sprintf("%s.go", object.Name))

	createFile(tmpls.MockBuilderTmpl, filename, object, destination)

}

func createObject(name string, fields string) Object {

	fieldsMapper := make(map[string]string)
	fieldsCloven := strings.Split(fields, ",")

	for _, f := range fieldsCloven {
		fieldsClovenByType := strings.Split(f, "-")
		if len(fieldsClovenByType) < 2 {
			log.Fatalf("There's one field who is not in Name-type")
		}
		key := fieldsClovenByType[0]
		value := fieldsClovenByType[1]
		v := checkSymbols(value)
		fieldsMapper[key] = v
	}

	return Object{
		Name:   name,
		Fields: fieldsMapper,
	}
}

func checkSymbols(value string) string {
	if strings.HasPrefix(value, "p") {
		return strings.Replace(value, "p", "*", -1)
	}
	if strings.HasPrefix(value, "a") {
		return strings.Replace(value, "a", "[]", -1)
	}
	return value
}

func createFile(MockBuilderTmpl string, outputFile string, data Object, destination string) {
	tmpl := template.Must(template.New("").
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

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.MkdirAll(destination, 0700)
	}

	path := filepath.Join(destination, outputFile)
	newFilePath := filepath.FromSlash(path)

	f, err := os.Create(newFilePath)
	if err != nil {
		log.Fatalf("Could not create file: %v\n", err)
	}

	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}
