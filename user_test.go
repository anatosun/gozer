package gozer

import (
	"testing"
)

func TestUser(t *testing.T) {
	client := Client()
	user, err := client.User(4055546082)
	if err != nil {
		t.Error(err)
	}
	if user.ID != 4055546082 {
		t.Errorf("Expected user ID to be 4055546082, got %d", user.ID)
	}
	if user.Name != "giligan06" {
		t.Errorf("Expected user name to be giligan06, got %s", user.Name)
	}

	if user.Type != "user" {
		t.Errorf("Expected user type to be user, got %s", user.Type)
	}
}
