package gozer

import (
	"os"
	"testing"
)

func TestAuthentication(t *testing.T) {

	client := Client()
	username := os.Getenv("DEEZER_USERNAME")
	password := os.Getenv("DEEZER_PASSWORD")

	if username == "" || password == "" {
		t.Skip("DEEZER_USERNAME or DEEZER_PASSWORD environment variables are not set")
	} else {

		err := client.authenticate(username, password)
		if err != nil {
			t.Fatal(err)
		}
	}

	err := client.authenticate("wrong@example.com", "wrong")
	if err == nil {
		t.Fatal("should raise an error")
	}

}
