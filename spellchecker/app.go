package spellchecker

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	endpoint = "/data/wordlist.txt"
	tmpFile  = "/tmp/sc_dictionary"
)

type App struct {
	HttpClient *http.Client
	BaseURL    *url.URL
}

func NewApp(baseURL string) (*App, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base URL: %w", err)
	}

	return &App{
		HttpClient: &http.Client{},
		BaseURL:    u,
	}, nil
}

func (a *App) Run() error {
	wl, _ := a.FetchWordList()

	c, err := itemCount(wl)
	if err != nil {
		return fmt.Errorf("unable to get item count: %w", err)
	}

	bf := &BloomFilter{
		ItemCount:                c,
		FalsePositiveProbability: 0.05,
	}

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
			return fmt.Errorf("unable to read input: %v", err)
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

func (a *App) FetchWordList() (io.ReadSeeker, error) {
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		b, err := a.fetchWordListFromRemote()

		if err != nil {
			return nil, fmt.Errorf("unable to fetch from remote: %v", err)
		}

		if err = ioutil.WriteFile(tmpFile, b, 0644); err != nil {
			return nil, fmt.Errorf("unable to write dictionary: %v", err)
		}

		return bytes.NewReader(b), nil
	}

	b, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read dictionary: %v", err)
	}

	return bytes.NewReader(b), nil
}

func (a *App) fetchWordListFromRemote() (wl []byte, e error) {
	rel := &url.URL{Path: endpoint}
	u := a.BaseURL.ResolveReference(rel)

	resp, err := a.HttpClient.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("unable to fetch word list: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			wl = nil
			e = fmt.Errorf("unable to close response body: %w", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unable to fetch word last status code: %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read word list: %w", err)
	}

	return b, nil
}

func itemCount(r io.ReadSeeker) (uint32, error) {
	var c uint32

	if err := seekStart(r); err != nil {
		return 0, err
	}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		c++
	}

	if err := seekStart(r); err != nil {
		return 0, err
	}

	return c, nil
}

func seekStart(s io.Seeker) error {
	_, err := s.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("unable to rewind word list: %w", err)
	}

	return nil
}
