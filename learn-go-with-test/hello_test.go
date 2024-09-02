package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Amol")
	want := "Hello Amol"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
