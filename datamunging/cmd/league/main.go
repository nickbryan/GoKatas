package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/nickbryan/GoKatas/datamunging/cmd/internal"
)

const path = "data/football.dat"

func main() {
	run(os.Stdout)
}

func run(w io.Writer) {
	c, err := readFile()
	if err != nil {
		log.Fatalln(err)
	}

	rxp := regexp.MustCompile("[^A-Z]*(.*\n)")
	input := rxp.ReplaceAllString(string(c), "$1")

	app := internal.NewApp(strings.NewReader(input), w, "Team", "F", "A")
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

func readFile() ([]byte, error) {
	_, p, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to gather Caller information")
	}

	b, err := ioutil.ReadFile(filepath.Join(filepath.Dir(p), path))
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	return b, nil
}
