package main

import (
	"bytes"
	"testing"
)

func Test_run(t *testing.T) {
	want := "Aston_Villa -- Spread: 1.000000\n"
	got := &bytes.Buffer{}
	run(got)
	if got.String() != want {
		t.Errorf("unexpected output: got: %s, want: %s", got, want)
	}
}
