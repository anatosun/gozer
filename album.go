package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type albums struct {
	Album []album `json:"data"`
	Total uint64  `json:"total"`
}

type album struct {
	ID                    uint64 `json:"id"`
	Title                 string `json:"title"`
	Upc                   string `json:"upc"`
	Link                  string `json:"link"`
	Share                 string `json:"share"`
	Cover                 string `json:"cover"`
	CoverSmall            string `json:"cover_small"`
	CoverMedium           string `json:"cover_medium"`
	CoverBig              string `json:"cover_big"`
	CoverXl               string `json:"cover_xl"`
	Md5Image              string `json:"md5_image"`
	GenreID               uint64 `json:"genre_id"`
	Genres                genres `json:"genres"`
	Label                 string `json:"label"`
	NbTracks              uint64 `json:"nb_tracks"`
	Duration              uint64 `json:"duration"`
	Fans                  uint64 `json:"fans"`
	ReleaseDate           string `json:"release_date"`
	RecordType            string `json:"record_type"`
	Available             bool   `json:"available"`
	Tracklist             string `json:"tracklist"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics uint64 `json:"explicit_content_lyrics"`
	ExplicitContentCover  uint64 `json:"explicit_content_cover"`
	Contributors          []struct {
		ID            uint64 `json:"id"`
		Name          string `json:"name"`
		Link          string `json:"link"`
		Share         string `json:"share"`
		Picture       string `json:"picture"`
		PictureSmall  string `json:"picture_small"`
		PictureMedium string `json:"picture_medium"`
		PictureBig    string `json:"picture_big"`
		PictureXl     string `json:"picture_xl"`
		Radio         bool   `json:"radio"`
		Tracklist     string `json:"tracklist"`
		Type          string `json:"type"`
		Role          string `json:"role"`
	} `json:"contributors"`
	Artist artist `json:"artist"`
	Type   string `json:"type"`
	Tracks tracks `json:"tracks"`
}

func (client *deezer) Album(id uint64) (album *album, err error) {
	url := fmt.Sprintf("%s/album/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return album, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return album, err
	}
	err = json.Unmarshal(bodyBytes, &album)
	if err != nil {
		return album, newDeezerError(bodyBytes)
	}

	return album, err
}
