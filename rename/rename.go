package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var name = flag.String("name", "", "New name for your app")

func main() {
	flag.Parse()

	if *name == "" {
		fmt.Println("No name given. Aborting.")

		return
	}

	err := filepath.Walk(".", handleFile)

	if err != nil {
		panic(err)
	}
}

func handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if filepath.Ext(path) == ".go" && filepath.Base(path) != "rename.go" {
		rwErr := rewriteImportsForFile(path)

		if rwErr != nil {
			return rwErr
		}
	}

	return nil
}

func rewriteImportsForFile(path string) error {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)

	if err != nil {
		return err
	}

	for _, imp := range file.Imports {
		rewrittenImport := strings.Replace(imp.Path.Value, `"bones/`, fmt.Sprintf(`"%s/`, *name), -1)
		imp.Path.Value = rewrittenImport
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, file)

	err = ioutil.WriteFile(path, buf.Bytes(), 0)

	if err != nil {
		return err
	}

	fmt.Printf("Updated imports for %s", path)

	return nil
}
