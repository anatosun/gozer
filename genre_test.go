package gozer

import (
	"testing"
)

func TestGenres(t *testing.T) {

	deezer := Client()
	genres, err := deezer.Genres()

	if err != nil {
		t.Error(err)
	}

	if len(genres.Genre) == 0 {
		t.Errorf("Expected genres to be greater than 0, got %d", len(genres.Genre))
	}

}

func TestGenre(t *testing.T) {

	deezer := Client()
	genre, err := deezer.Genre(132)

	if err != nil {
		t.Fatal(err)
	}

	if genre.ID != 132 {
		t.Errorf("Expected genre ID to be 132, got %d", genre.ID)
	}

	if genre.Name != "Pop" {
		t.Errorf("Expected genre name to be Pop, got %s", genre.Name)
	}

	if genre.Type != "genre" {
		t.Errorf("Expected genre type to be genre, got %s", genre.Type)
	}

}
