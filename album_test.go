package gozer

import (
	"testing"
)

func TestAlbum(t *testing.T) {
	deezer := Client()
	var ID uint64
	ID = 302127
	album, err := deezer.Album(ID)
	if err != nil {
		t.Fatal(err)
	}

	if album.ID != ID {

		t.Fatalf("Expected %d, got %d", ID, album.ID)

	}

	if album.Title != "Discovery" {
		t.Fatalf("Expected %s, got %s", "Discovery", album.Title)
	}

	if album.Artist.Name != "Daft Punk" {
		t.Fatalf("Expected %s, got %s", "Daft Punk", album.Artist.Name)
	}

	if album.Tracks.Track[0].Title != "One More Time" {
		t.Fatalf("Expected %s, got %s", "One More Time", album.Tracks.Track[0].Title)
	}

	if album.Tracks.Track[1].Title != "Aerodynamic" {
		t.Fatalf("Expected %s, got %s", "Aerodynamic", album.Tracks.Track[1].Title)
	}

	if album.Tracks.Track[2].Title != "Digital Love" {
		t.Fatalf("Expected %s, got %s", "Digital Love", album.Tracks.Track[2].Title)
	}

	if album.Tracks.Track[3].Title != "Harder, Better, Faster, Stronger" {
		t.Fatalf("Expected %s, got %s", "Harder, Better, Faster, Stronger", album.Tracks.Track[3].Title)
	}

	if album.Tracks.Track[4].Title != "Crescendolls" {
		t.Fatalf("Expected %s, got %s", "Crescendolls", album.Tracks.Track[4].Title)
	}

	if album.Tracks.Track[5].Title != "Nightvision" {
		t.Fatalf("Expected %s, got %s", "Nightvision", album.Tracks.Track[5].Title)
	}

	if album.Tracks.Track[6].Title != "Superheroes" {
		t.Fatalf("Expected %s, got %s", "Superheroes", album.Tracks.Track[6].Title)
	}

	if album.Tracks.Track[7].Title != "High Life" {
		t.Fatalf("Expected %s, got %s", "High Life", album.Tracks.Track[7].Title)
	}

	if album.Tracks.Track[8].Title != "Something About Us" {
		t.Fatalf("Expected %s, got %s", "Something About Us", album.Tracks.Track[8].Title)
	}

	if album.Tracks.Track[9].Title != "Voyager" {
		t.Fatalf("Expected %s, got %s", "Voyager", album.Tracks.Track[9].Title)
	}

	if album.Tracks.Track[10].Title != "Veridis Quo" {
		t.Fatalf("Expected %s, got %s", "Veridis Quo", album.Tracks.Track[10].Title)
	}

	if album.Tracks.Track[11].Title != "Short Circuit" {
		t.Fatalf("Expected %s, got %s", "Short Circuit", album.Tracks.Track[11].Title)
	}

	if album.Tracks.Track[12].Title != "Face to Face" {
		t.Fatalf("Expected %s, got %s", "Face to Face", album.Tracks.Track[12].Title)
	}

	if album.Tracks.Track[13].Title != "Too Long" {
		t.Fatalf("Expected %s, got %s", "Too Long", album.Tracks.Track[13].Title)
	}

}
