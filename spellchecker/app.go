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

// App encapsulates the functionality of the spellchecker.
type App struct {
	HttpClient *http.Client
	BaseURL    *url.URL
	Output     io.Writer
	Input      io.Reader
}

// NewApp will create a new app and parse the baseURL.
func NewApp(baseURL string) (*App, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base URL: %w", err)
	}

	return &App{
		HttpClient: &http.Client{},
		BaseURL:    u,
		Output:     os.Stdout,
		Input:      os.Stdin,
	}, nil
}

// Run will load the dictionary into the bloom filter before starting a loop to listen
// for user input. It will check the input from the user against the bloom filter to see if
// the worlds are correct according to the dictionary.
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

	r := bufio.NewReader(a.Input)
	for {
		if _, err := a.Output.Write([]byte("Check word: ")); err != nil {
			return fmt.Errorf("unable to write output: %w", err)
		}

		i, err := r.ReadString('\n')
		if err != nil {
			return fmt.Errorf("unable to read input: %v", err)
		}

		i = strings.Trim(i, "\n")

		if i == "q" {
			break
		}

		var output string
		if bf.Exists(i) {
			output = fmt.Sprintf("%s is spelled correctly.\n", i)
		} else {
			output = fmt.Sprintf("%s is not spelled correctly.\n", i)
		}

		if _, err := a.Output.Write([]byte(output)); err != nil {
			return fmt.Errorf("unable to write output: %w", err)
		}
	}

	return nil
}

// FetchWorldList will check for a tmp file before fetching the fill dictionary from the remote service.
// If the tmp file exists then that will be used to save downloading the full dictionary. If the dictionary
// has to be downloaded then it will be written to a tmp file to make future runs faster.
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
