package entities

import (
	"testing"
)

func TestUserEquals(t *testing.T) {
	a := []User{
		User{Id: 1, Email: "a@foo.com"},
		User{Id: 2, Email: "b@foo.com"},
	}

	b := []User{
		User{Id: 1, Email: "a@foo.com"},
		User{Id: 3, Email: "c@foo.com"},
	}

	c := []User{
		User{Id: 1, Email: "a@foo.com"},
		User{Id: 2, Email: "b@foo.com"},
	}

	if UserEquals(a, b) {
		t.Error("Expected false")
	}

	if !UserEquals(a, c) {
		t.Error("Expected true")
	}
}
