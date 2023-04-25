package gozer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type infos struct {
	CountryIso          string `json:"country_iso"`
	Country             string `json:"country"`
	Open                bool   `json:"open"`
	Pop                 string `json:"pop"`
	UploadToken         string `json:"upload_token"`
	UploadTokenLifetime uint64 `json:"upload_token_lifetime"`
	UserToken           string `json:"user_token"`
	Hosts               struct {
		Stream string `json:"stream"`
		Images string `json:"images"`
	} `json:"hosts"`
	Ads struct {
		Audio struct {
			Default struct {
				Start    uint64 `json:"start"`
				Interval uint64 `json:"interval"`
				Unit     string `json:"unit"`
			} `json:"default"`
		} `json:"audio"`
		Display struct {
			Interstitial struct {
				Start    uint64 `json:"start"`
				Interval uint64 `json:"interval"`
				Unit     string `json:"unit"`
			} `json:"interstitial"`
		} `json:"display"`
		BigNativeAdsHome struct {
			Iphone struct {
				Enabled bool `json:"enabled"`
			} `json:"iphone"`
			Ipad struct {
				Enabled bool `json:"enabled"`
			} `json:"ipad"`
			Android struct {
				Enabled bool `json:"enabled"`
			} `json:"android"`
			AndroidTablet struct {
				Enabled bool `json:"enabled"`
			} `json:"android_tablet"`
		} `json:"big_native_ads_home"`
	} `json:"ads"`
	HasPodcasts       bool          `json:"has_podcasts"`
	Offers            []interface{} `json:"offers"`
	FreemiumAvailable bool          `json:"freemium_available"`
	LyricsAvailable   bool          `json:"lyrics_available"`
}

func (client *deezer) Infos() (infos *infos, err error) {
	url := fmt.Sprintf("%s/genre", client.base)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return infos, err
	}
	err = json.Unmarshal(bodyBytes, &infos)

	if err != nil {
		return infos, err
	}

	return nil, newDeezerError(bodyBytes)
}
