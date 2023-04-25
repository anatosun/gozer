package gozer

import (
	"testing"
)

func TestArtist(t *testing.T) {

	deezer := Client()
	artist, err := deezer.Artist(27)

	if err != nil {
		t.Error(err)
	}

	if artist.ID != 27 {
		t.Errorf("Expected artist ID to be 27, got %d", artist.ID)
	}

	if artist.Name != "Daft Punk" {
		t.Errorf("Expected artist name to be Daft Punk, got %s", artist.Name)
	}

	if artist.Type != "artist" {
		t.Errorf("Expected artist type to be artist, got %s", artist.Type)
	}

	if artist.Radio != true {
		t.Errorf("Expected artist radio to be true, got %t", artist.Radio)
	}

}
