package spellchecker

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	endpoint = "/data/wordlist.txt"
	tmpFile  = "/tmp/sc_dictionary"
)

type App struct {
	HttpClient *http.Client
	BaseURL    *url.URL
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
