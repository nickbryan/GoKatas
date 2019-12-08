package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nickbryan/GoKatas/datamunging/cmd/internal"
)

const path = "/data/weather.dat"

func main() {
	run(os.Stdout)
}

func run(w io.Writer) {
	f, err := readFile()
	if err != nil {
		log.Fatalln(err)
	}

	app := internal.NewApp(f, w, "Dy", "MnT", "MxT")
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

func readFile() (*os.File, error) {
	_, p, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to gather Caller information")
	}

	f, err := os.Open(filepath.Join(filepath.Dir(p), path))
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	return f, nil
}
