package main

import (
	"bytes"
	"testing"
)

func Test_run(t *testing.T) {
	want := "14 -- Spread: 2.000000\n"
	got := &bytes.Buffer{}
	run(got)
	if got.String() != want {
		t.Errorf("unexpected output: got: %s, want: %s", got, want)
	}
}
