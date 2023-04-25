package gozer

import "testing"

func TestTrac(t *testing.T) {
	client := Client()

	track, err := client.Track(3135556)
	if err != nil {
		t.Error(err)
	}

	if track.ID != 3135556 {
		t.Fatalf("Expected track ID to be 3135556, got %d", track.ID)

	}

	if track.Title != "Harder, Better, Faster, Stronger" {
		t.Fatalf("Expected track Title to be Harder, Better, Faster, Stronger, got %s", track.Title)
	}

	if track.TitleShort != "Harder, Better, Faster, Stronger" {
		t.Fatalf("Expected track TitleShort to be Harder, Better, Faster, Stronger, got %s", track.TitleShort)
	}
	if track.TitleVersion != "" {
		t.Fatalf("Expected track TitleVersion to be empty, got %s", track.TitleVersion)
	}

	if track.Isrc != "GBDUW0000059" {
		t.Fatalf("Expected track Isrc to be GBDUW0000059, got %s", track.Isrc)
	}

	if track.Duration != 224 {
		t.Fatalf("Expected track Duration to be 224, got %d", track.Duration)
	}

	if track.TrackPosition != 4 {
		t.Fatalf("Expected track TrackPosition to be 4, got %d", track.TrackPosition)
	}

	if track.DiskNumber != 1 {
		t.Fatalf("Expected track DiskNumber to be 1, got %d", track.DiskNumber)
	}

	if track.Rank != 778628 {
		t.Fatalf("Expected track Rank to be 778628, got %d", track.Rank)
	}

	if track.ReleaseDate != "2005-01-24" {
		t.Fatalf("Expected track ReleaseDate to be 2005-01-24, got %s", track.ReleaseDate)
	}

	if track.ExplicitLyrics != false {
		t.Fatalf("Expected track ExplicitLyrics to be false, got %t", track.ExplicitLyrics)
	}

	if track.ExplicitContentLyrics != 0 {
		t.Fatalf("Expected track ExplicitContentLyrics to be 0, got %d", track.ExplicitContentLyrics)
	}

	if track.Bpm != 123.4 {
		t.Fatalf("Expected track Bpm to be 123.4, got %f", track.Bpm)
	}

}
