package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type chart struct {
	Tracks    tracks    `json:"tracks"`
	Albums    albums    `json:"albums"`
	Artists   artists   `json:"artists"`
	Playlists playlists `json:"playlists"`
	Podcasts  podcasts  `json:"podcasts"`
}

func (client *deezer) Chart(id uint64) (chart *chart, err error) {
	url := fmt.Sprintf("%s/chart/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return chart, err
	}
	err = json.Unmarshal(bodyBytes, &chart)

	if err != nil {
		return chart, err
	}
	if len(chart.Tracks.Track) == 0 {
		return chart, newDeezerError(bodyBytes)
	}

	return chart, err
}
