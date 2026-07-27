package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Amol")

	got := buffer.String()
	want := "Hello, Amol"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
