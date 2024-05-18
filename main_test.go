package main

import (
	"github.com/martinheinrich2/gowordcount/cmd"
	"testing"
)

func TestCountBytes(t *testing.T) {
	got := cmd.CountBytes("test.txt")
	want := 2780

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountLines(t *testing.T) {
	got := cmd.CountLines("test.txt")
	want := 33

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountWords(t *testing.T) {
	got := cmd.CountWords("test.txt")
	want := 409

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountCharacters(t *testing.T) {
	got := cmd.CountCharacters("test.txt")
	want := 2780
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
