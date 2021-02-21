package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sarath")
	want := "Hello, Sarath"
	if got != want {
		t.Errorf("got %q , want %q", got, want)
	}
}
