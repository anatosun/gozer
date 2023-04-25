package gozer

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type grant struct {
	AccessToken    string `json:"access_token"`
	Expire         uint64 `json:"expire"`
	Expires        uint64 `json:"expires"`
	IsMultiaccount bool   `json:"is_multiaccount"`
}
type arl struct {
	Error   []*DeezerError `json:"error"`
	Results string         `json:"results"`
}

type deezer struct {
	base   string
	id     string
	secret string
	grant  *grant
	ARL    string
}

func Client() (client *deezer) {

	client = &deezer{}
	client.base = "https://api.deezer.com"
	client.id = "172365"
	client.secret = "fb0bec7ccc063dab0417eb7b0d847f34"
	return client
}

func (client *deezer) authenticate(email string, password string) (err error) {
	data := md5.Sum([]byte(password))
	hashedPassword := fmt.Sprintf("%x", data)
	accessString := client.id + email + hashedPassword + client.secret
	hash := md5.Sum([]byte(accessString))

	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) "
	reqUrl := fmt.Sprintf("%s/auth/token", client.base)

	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", userAgent)
	req.URL.RawQuery = url.Values{
		"app_id":   {client.id},
		"login":    {email},
		"password": {hashedPassword},
		"hash":     {fmt.Sprintf("%x", hash)},
	}.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	grant := &grant{}
	err = json.Unmarshal(bodyBytes, grant)
	if err != nil {
		return err
	}

	if grant.AccessToken == "" {
		return newDeezerError(bodyBytes)
	}

	client.grant = grant

	req, err = http.NewRequest("GET", "https://api.deezer.com/platform/generic/track/3135556", nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.grant.AccessToken))
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	req, err = http.NewRequest("GET", "https://www.deezer.com/ajax/gw-light.php?method=user.getArl&input=3&api_version=1.0&api_token=null", nil)
	req.Header.Set("User-Agent", userAgent)
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("Authorization", "Bearer "+client.grant.AccessToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	arl := &arl{}
	err = json.Unmarshal(bodyBytes, arl)
	if err != nil {
		return err
	}

	if len(arl.Error) > 0 {
		return arl.Error[0]
	}

	if arl.Results == "" {
		return newDeezerError(bodyBytes)
	}

	client.ARL = arl.Results

	return nil
}
