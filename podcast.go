package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type podcasts struct {
	Podcast []podcast `json:"data"`
	Total   uint64    `json:"total"`
}

type podcast struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Available     bool   `json:"available"`
	Fans          uint64 `json:"fans"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

func (client *deezer) Podcast(id uint64) (podcast *podcast, err error) {
	url := fmt.Sprintf("%s/podcast/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return podcast, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return podcast, err
	}
	err = json.Unmarshal(bodyBytes, &podcast)
	if err != nil {
		return podcast, err
	}
	if podcast.ID == 0 {
		return podcast, newDeezerError(bodyBytes)
	}

	return podcast, err
}
