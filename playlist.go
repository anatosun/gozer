package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type playlists struct {
	Playlists []playlist `json:"data"`
	Total     uint64     `json:"total"`
}

type playlist struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Duration      int    `json:"duration"`
	Public        bool   `json:"public"`
	IsLovedTrack  bool   `json:"is_loved_track"`
	Collaborative bool   `json:"collaborative"`
	NbTracks      int    `json:"nb_tracks"`
	Fans          int    `json:"fans"`
	Link          string `json:"link"`
	Share         string `json:"share"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Checksum      string `json:"checksum"`
	Tracklist     string `json:"tracklist"`
	CreationDate  string `json:"creation_date"`
	Md5Image      string `json:"md5_image"`
	PictureType   string `json:"picture_type"`
	Creator       struct {
		ID        uint64 `json:"id"`
		Name      string `json:"name"`
		Tracklist string `json:"tracklist"`
		Type      string `json:"type"`
	} `json:"creator"`
	Type   string `json:"type"`
	Tracks tracks `json:"tracks"`
}

func (client *deezer) Playlist(id uint64) (playlist *playlist, err error) {
	url := fmt.Sprintf("%s/playlist/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return playlist, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return playlist, err
	}
	err = json.Unmarshal(bodyBytes, &playlist)
	if err != nil {
		return playlist, err
	}

	if len(playlist.Tracks.Track) == 0 {
		return playlist, newDeezerError(bodyBytes)
	}
	return playlist, err
}
