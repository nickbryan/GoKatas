package main

import (
	"flag"
	"log"
	"os"

	"github.com/nickbryan/GoKatas/datamunging"
)

func main() {
	var err error

	file := flag.String("file", "", "the path to the data file")
	flag.Parse()

	if *file == "" {
		log.Fatal("file not specified")
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	rows, err := datamunging.Parse(f)
	if err != nil {
		log.Fatalf("unable to parse file: %v", err)
	}

	if err = datamunging.WriteMinSpread(&rows, os.Stdout); err != nil {
		log.Fatalf("unable to write to Stdout: %v", err)
	}

}
