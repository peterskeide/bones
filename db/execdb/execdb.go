package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	initFile string = "database.init"
)

var connStr = flag.String("c", "", "String to use when connecting to database")
var filename = flag.String("f", initFile, "Name of sql file to execute, relative to script directory. If not given, all the scripts listed in database.init will be executed")
var scriptDir = flag.String("d", "./db/scripts", "Name of directory with sql scripts")
var filenames []string

func main() {
	flag.Parse()
	initFilenames(*filename)
	withTx(execFiles)
}

func initFilenames(filename string) {
	if filename == initFile {
		getFilenamesFromInitFile()
	} else {
		filenames = []string{filename}
	}
}

func getFilenamesFromInitFile() {
	content, err := ioutil.ReadFile(filepath.Join(*scriptDir, initFile))

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line != "" {
			filenames = append(filenames, strings.TrimSpace(line))
		}
	}
}

func withTx(fn func(*sql.Tx) error) {
	db, err := sql.Open("postgres", *connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	err = fn(tx)

	if err != nil {
		tx.Rollback()

		panic(err)
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}

func execFiles(tx *sql.Tx) error {
	for _, basename := range filenames {
		filename := filepath.Join(*scriptDir, basename)

		err := execFile(filename, tx)

		if err != nil {
			return err
		}
	}

	return nil
}

func execFile(filename string, tx *sql.Tx) error {
	cmd, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	_, err = tx.Exec(string(cmd))

	if err != nil {
		return err
	}

	fmt.Println(filename)

	return nil
}
