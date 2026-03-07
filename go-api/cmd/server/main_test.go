package main

import (
	"testing"
)

func TestServerPort_Default(t *testing.T) {
	t.Setenv("SERVER_PORT", "")

	got := serverPort()
	if got != "8080" {
		t.Fatalf("expected default port 8080, got %q", got)
	}
}

func TestServerPort_FromEnv(t *testing.T) {
	t.Setenv("SERVER_PORT", "9090")

	got := serverPort()
	if got != "9090" {
		t.Fatalf("expected port 9090, got %q", got)
	}
}
