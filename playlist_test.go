package gozer

import (
	"testing"
)

func TestPlaylist(t *testing.T) {
	client := Client()
	playlist, err := client.Playlist(11311474964)
	if err != nil {
		t.Error(err)
	}

	if playlist.Title != "Example" {
		t.Error("Playlist title is not correct")
	}

	if playlist.Description != "" {
		t.Error("Playlist description is not correct")
	}

	if playlist.Duration != 21810 {
		t.Fatalf("Playlist duration is not correct, got %d, expected %d", playlist.Duration, 21810)
	}

	if playlist.NbTracks != 50 {

		t.Fatalf("Playlist nbTracks is not correct, got %d, expected %d", playlist.NbTracks, 50)
	}

}
