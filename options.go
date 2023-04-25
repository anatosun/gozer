package gozer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Options struct {
	Streaming         bool   `json:"streaming"`
	StreamingDuration uint64 `json:"streaming_duration"`
	Offline           bool   `json:"offline"`
	Hq                bool   `json:"hq"`
	AdsDisplay        bool   `json:"ads_display"`
	AdsAudio          bool   `json:"ads_audio"`
	TooManyDevices    bool   `json:"too_many_devices"`
	CanSubscribe      bool   `json:"can_subscribe"`
	RadioSkips        uint64 `json:"radio_skips"`
	Lossless          bool   `json:"lossless"`
	Preview           bool   `json:"preview"`
	Radio             bool   `json:"radio"`
	Type              string `json:"type"`
}

func (client *deezer) GetOptions() (options *Options, err error) {
	url := client.base + "/editorial"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return options, err
	}
	err = json.Unmarshal(bodyBytes, &options)

	if err != nil {
		return options, err
	}

	return options, newDeezerError(bodyBytes)

}
