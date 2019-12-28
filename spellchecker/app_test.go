package spellchecker

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

type fetchTC struct {
	payload    string
	statusCode int
	err        error
	tmpDic     string
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
		"successful payload is returned": {
			statusCode: http.StatusOK,
			payload:    "hello\nworld",
			err:        nil,
			tmpDic:     "",
		},
		"unsuccessful status code is handled": {
			statusCode: http.StatusNotFound,
			err:        errors.New("unable to fetch from remote: unable to fetch word last status code: 404"),
			tmpDic:     "",
		},
		"tmp file is not updated if exists": {
			payload: "hello\nworld\n",
			tmpDic:  "hello\nworld\n",
		},
	}

	cleanTmpFile(t)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.tmpDic != "" {
				writeTmpFile(t, tc.tmpDic)
			}
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
					t.Errorf("incorrect retrun: expected: %s, got: %s", tc.payload, b.String())
				}
				if readTmpFile(t) != tc.payload {
					t.Errorf("tmp file wrong contents: expected: %s, got: %s", tc.payload, b.String())
				}
			}
		})
	}
}

func TestApp_Run(t *testing.T) {
	testDic := `
hello
world
this
is
a
test
dictionary
that
does
not
contain
all
of
the
world
`

	tests := map[string]struct {
		dic    string
		input  string
		output string
	}{
		"world is correct if in list":             {dic: testDic, input: "hello\nq\n", output: "Check word: hello is spelled correctly.\nCheck word: "},
		"world is correct if in list second word": {dic: testDic, input: "world\nq\n", output: "Check word: world is spelled correctly.\nCheck word: "},
		"world is incorrect if not in list":       {dic: testDic, input: "pizza\nq\n", output: "Check word: pizza is not spelled correctly.\nCheck word: "},
	}

	cleanTmpFile(t)
	s := httptest.NewServer(nil)
	defer s.Close()

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer cleanTmpFile(t)
			writeTmpFile(t, tc.dic)

			got := &bytes.Buffer{}
			w := strings.NewReader(tc.input)
			a := &App{HttpClient: s.Client(), BaseURL: new(url.URL), Output: got, Input: w}
			if err := a.Run(); err != nil {
				t.Fatalf("run caused an errror: %v", err)
			}

			if got.String() != tc.output {
				t.Errorf("expected: %s, got: %s", tc.output, got.String())
			}
		})
	}
}

func writeTmpFile(t *testing.T, s string) {
	if err := ioutil.WriteFile(tmpFile, []byte(s), 0644); err != nil {
		t.Fatalf("unable to write dictionary: %v", err)
	}
}

func readTmpFile(t *testing.T) string {
	b, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("unable to read dictionary: %v", err)
	}

	return string(b)
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
