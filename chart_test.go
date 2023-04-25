package gozer

import "testing"

func TestChart(t *testing.T) {

	client := Client()
	chart, err := client.Chart(0)
	if err != nil {
		t.Error(err)
	}

	if len(chart.Tracks.Track) == 0 {
		t.Error("Expected tracks to be non-empty")
	}

	if len(chart.Albums.Album) == 0 {
		t.Error("Expected albums to be non-empty")
	}

	if len(chart.Artists.Artist) == 0 {
		t.Error("Expected artists to be non-empty")
	}

}
