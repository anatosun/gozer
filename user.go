package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID                             uint64   `json:"id"`
	Name                           string   `json:"name"`
	Lastname                       string   `json:"lastname"`
	Firstname                      string   `json:"firstname"`
	Email                          string   `json:"email"`
	Status                         uint64   `json:"status"`
	Birthday                       string   `json:"birthday"`
	InscriptionDate                string   `json:"inscription_date"`
	Gender                         string   `json:"gender"`
	Link                           string   `json:"link"`
	Picture                        string   `json:"picture"`
	PictureSmall                   string   `json:"picture_small"`
	PictureMedium                  string   `json:"picture_medium"`
	PictureBig                     string   `json:"picture_big"`
	PictureXl                      string   `json:"picture_xl"`
	Country                        string   `json:"country"`
	Lang                           string   `json:"lang"`
	IsKid                          bool     `json:"is_kid"`
	ExplicitContentLevel           string   `json:"explicit_content_level"`
	ExplicitContentLevelsAvailable []string `json:"explicit_content_levels_available"`
	Tracklist                      string   `json:"tracklist"`
	IsMultiaccount                 bool     `json:"is_multiaccount"`
	IsMultiaccountParent           bool     `json:"is_multiaccount_parent"`
	Md5Image                       string   `json:"md5_image"`
	IsFlowAvailable                bool     `json:"is_flow_available"`
	Type                           string   `json:"type"`
}

func (client *deezer) User(id uint64) (user *user, err error) {
	url := fmt.Sprintf("%s/user/%d", client.base, id)
	resp, err := http.Get(url)
	if err != nil {
		return user, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, newDeezerError(bodyBytes)
	}

	return user, err
}
