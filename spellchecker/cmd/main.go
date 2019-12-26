package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/nickbryan/GoKatas/spellchecker"
)

const baseURL = "http://codekata.com/"

func main() {
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("unable to parse URL: %v", err))
	}

	app := &spellchecker.App{
		HttpClient: &http.Client{},
		BaseURL:    u,
	}
	wl, err := app.FetchWordList()
	if err != nil {
		log.Fatalln(fmt.Sprintf("unable to fetch word list: %v", err))
	}

	bf := &spellchecker.BloomFilter{
		ItemCount:                itemCount(wl),
		FalsePositiveProbability: 0.05,
	}

	wl.Seek(0, io.SeekStart)
	s := bufio.NewScanner(wl)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		bf.Add(s.Text())
	}

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Check word: ")
		i, err := r.ReadString('\n')
		if err != nil {
			log.Fatalln(fmt.Sprintf("unable to read input: %v", err))
		}

		i = strings.Trim(i, "\n")

		var output string
		if bf.Exists(i) {
			output = fmt.Sprintf("%s is spelled correctly.", i)
		} else {
			output = fmt.Sprintf("%s is not spelled correctly.", i)
		}
		fmt.Println(output)
	}

}

func itemCount(r io.ReadSeeker) uint32 {
	var c uint32
	r.Seek(0, io.SeekStart)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		c++
	}

	return c
}
