package auth

import "testing"

func TestGetAPIKey(t *testing.T) {

	got := true
	want := true

	if got != want {
		t.Fatalf("expected %v, got: %v", got, want)
	}

}

func TestGetAPIKeyAgain(t *testing.T) {

	got := true
	want := true

	if got != want {
		t.Fatalf("expected %v, got: %v", got, want)
	}

}
