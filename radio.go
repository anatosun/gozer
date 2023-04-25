package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type radio struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Tracklist     string `json:"tracklist"`
	Md5Image      string `json:"md5_image"`
	Type          string `json:"type"`
}
type radios struct {
	Radio []radio `json:"data"`
}

func (client *deezer) Radios() (radios *radios, err error) {
	url := client.base + "/radio"
	resp, err := http.Get(url)
	if err != nil {
		return radios, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return radios, err
	}
	err = json.Unmarshal(bodyBytes, &radios)

	if err != nil {
		return radios, err
	}
	if len(radios.Radio) == 0 {
		return radios, newDeezerError(bodyBytes)
	}

	return radios, err
}
func (client *deezer) Radio(id uint64) (radio *radio, err error) {
	url := fmt.Sprintf("%s/radio/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return radio, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return radio, err
	}
	err = json.Unmarshal(bodyBytes, &radio)

	if err != nil {
		return radio, err
	}
	if radio.ID == 0 {
		return radio, newDeezerError(bodyBytes)
	}

	return radio, err
}
