package internal

import "testing"

func TestEmptyAuth(t *testing.T) {
	a := auth{
		username: " ",
		password: " ",
	}

	valid := ""
	invalid := "take:me"

	if !a.authenticate(valid) {
		t.Error("failed to check empty token")
	}

	if !a.authenticate(invalid) {
		t.Error("failed to check free token")
	}
}

func TestAuth(t *testing.T) {
	a := auth{
		username: "root",
		password: "password",
	}

	valid := "root:password"
	invalid := "take:me"

	if !a.authenticate(valid) {
		t.Error("failed to check true token")
	}

	if a.authenticate(invalid) {
		t.Error("failed to check free token")
	}
}
