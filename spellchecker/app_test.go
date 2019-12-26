package spellchecker

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

type fetchTC struct {
	payload    string
	statusCode int
	err        error
	t          *testing.T
}

func (ftc *fetchTC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != endpoint {
		ftc.t.Errorf("url path not expected: want: %v, got: %v", endpoint, r.URL.Path)
	}

	w.WriteHeader(ftc.statusCode)

	if _, err := w.Write([]byte(ftc.payload)); err != nil {
		ftc.t.Fatalf("unabke to write response payload: %v", err)
	}
}

func TestApp_FetchWordList(t *testing.T) {
	tests := map[string]*fetchTC{
		"successful payload is returned": &fetchTC{
			statusCode: http.StatusOK,
			payload:    "hello\nworld",
			err:        nil,
		},
		"unsuccessful status code is handled": &fetchTC{
			statusCode: http.StatusNotFound,
			err:        errors.New("unable to fetch from remote: unable to fetch word last status code: 404"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer cleanTmpFile(t)

			tc.t = t

			s := httptest.NewServer(tc)
			defer s.Close()

			u, err := url.Parse(s.URL)
			if err != nil {
				t.Fatalf("unable to parse URL: %v", err)
			}

			a := &App{HttpClient: s.Client(), BaseURL: u}

			got, err := a.FetchWordList()
			if err != nil && (tc.err == nil || tc.err.Error() != err.Error()) {
				t.Fatalf("error not expected: want: %v, got: %v", tc.err, err)
			}

			if err == nil {
				b := new(bytes.Buffer)
				if _, err := b.ReadFrom(got); err != nil {
					t.Fatalf("unable to read from response: %v", err)
				}
				if b.String() != tc.payload {
					t.Errorf("expected: %s, got: %s", tc.payload, b.String())
				}
			}
		})
	}
}

func cleanTmpFile(t *testing.T) {
	_, err := os.Stat(tmpFile)
	if os.IsNotExist(err) {
		return
	}

	if err := os.Remove(tmpFile); err != nil {
		t.Fatalf("unable to remove tmp file: %v", err)
	}
}
