package gozer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Editorial struct {
	Data []struct {
		ID            uint64 `json:"id"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`
		Type          string `json:"type"`
	} `json:"data"`
	Total uint64 `json:"total"`
}

func (client *deezer) GetEditorial() (editorial *Editorial, err error) {
	url := client.base + "/editorial"
	resp, err := http.Get(url)
	if err != nil {
		return editorial, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return editorial, err
	}
	err = json.Unmarshal(bodyBytes, &editorial)

	if err != nil {
		return editorial, err
	}
	if len(editorial.Data) == 0 {
		return editorial, newDeezerError(bodyBytes)
	}

	return editorial, err
}
