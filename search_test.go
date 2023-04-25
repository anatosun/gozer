package gozer

import "testing"

func TestSearch(t *testing.T) {
	deezer := Client()
	_, err := deezer.Search("eminem")
	if err != nil {
		t.Fatal(err)
	}

}
