package gozer

import (
	"testing"
)

func TestRadios(t *testing.T) {
	client := Client()
	radios, err := client.Radios()
	if err != nil {
		t.Error(err)
	}
	if len(radios.Radio) == 0 {
		t.Fatalf("Expected radios to be greater than 0, got %d", len(radios.Radio))
	}

}

func TestRadio(t *testing.T) {
	client := Client()
	radio, err := client.Radio(37151)

	if err != nil {
		t.Fatal(err)
	}

	if radio.ID != 37151 {
		t.Errorf("Expected radio ID to be 37151, got %d", radio.ID)
	}

	if radio.Title != "Hits" {
		t.Errorf("Expected radio title to be Hits, got %s", radio.Title)
	}

	if radio.Description != "Hits" {
		t.Errorf("Expected radio description to be Hits, got %s", radio.Description)
	}

	if radio.Type != "radio" {
		t.Errorf("Expected radio type to be radio, got %s", radio.Type)
	}

}
