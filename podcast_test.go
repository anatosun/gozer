package gozer

import "testing"

func TestPodcast(t *testing.T) {
	client := Client()
	podcast, err := client.Podcast(3823417)
	if err != nil {
		t.Error(err)
	}
	if podcast.ID != 3823417 {
		t.Fatalf("Podcast ID is %d, expected 3823417", podcast.ID)
	}

	if podcast.Title != "Die Nervigen" {
		t.Fatalf("Podcast title is %s, expected Die Nervigen", podcast.Title)
	}

	if podcast.Type != "podcast" {
		t.Fatalf("Podcast type is %s, expected podcast", podcast.Type)
	}

}
