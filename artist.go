package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type artists struct {
	Artist []artist `json:"data"`
	Total  uint64   `json:"total"`
}

type artist struct {
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	NbAlbum       uint64 `json:"nb_album"`
	NbFan         uint64 `json:"nb_fan"`
	Radio         bool   `json:"radio"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
}

func (client *deezer) Artist(id uint64) (artist *artist, err error) {
	url := fmt.Sprintf("%s/artist/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return artist, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return artist, err
	}
	err = json.Unmarshal(bodyBytes, &artist)

	if err != nil {
		return artist, err
	}

	if artist.ID == 0 {
		return artist, newDeezerError(bodyBytes)
	}

	return artist, err
}
