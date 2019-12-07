package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nickbryan/GoKatas/datamunging/cmd/internal"
)

const path = "/data/weather.dat"

func main() {
	f, err := readFile()
	if err != nil {
		log.Fatalln(err)
	}

	app := internal.NewApp(f, os.Stdout, "Dy", "MnT", "MxT")
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}

func readFile() (*os.File, error) {
	_, p, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("unable to gather Caller information")
	}

	f, err := os.Open(filepath.Dir(p) + path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	return f, nil
}
