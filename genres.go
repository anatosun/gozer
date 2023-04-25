package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type genre struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

type genres struct {
	Genre []genre `json:"data"`
}

func (client *deezer) Genres() (genres *genres, err error) {
	url := fmt.Sprintf("%s/genre", client.base)
	resp, err := http.Get(url)
	if err != nil {
		return genres, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return genres, err
	}
	err = json.Unmarshal(bodyBytes, &genres)
	if err != nil {
		return genres, err
	}
	if len(genres.Genre) == 0 {
		return genres, newDeezerError(bodyBytes)
	}

	return genres, err
}

func (client *deezer) Genre(id uint64) (genre *genre, err error) {
	url := fmt.Sprintf("%s/genre/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return genre, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return genre, err
	}
	err = json.Unmarshal(bodyBytes, &genre)
	if err != nil {
		return genre, err
	}
	if genre.ID == 0 {
		return genre, newDeezerError(bodyBytes)
	}

	return genre, err
}
