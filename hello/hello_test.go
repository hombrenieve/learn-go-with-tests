package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pablo")
	want := "Hello Pablo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
