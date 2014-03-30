package tictail

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
		AUTH_URL = "https://tictail.com/oauth/token"
)

type SampleTictailAuth struct {
	clientId string
	clientSecret string
	code string
	accessToken string
}

type TAuth struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	Store map[string]interface{} `json:"store"`
}

func newSampleTictailAuth(clientId, clientSecret, code string) (SampleTictailAuth, error){
	sta := SampleTictailAuth{
		clientId,
		clientSecret,
		code,
		"",
	}
	err := sta.getToken()
	return sta, err
}

func (sta SampleTictailAuth) getToken() error {
	resp, err := http.PostForm(AUTH_URL,
		url.Values{
			"client_id": {sta.clientId},
			"client_secret": {sta.clientSecret},
			"code": {sta.code},
			"grant_type": {"authorization_code"},
		});
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var tauth TAuth 
	err = json.Unmarshal(body, tauth)
	sta.accessToken = tauth.AccessToken
	return err
}

func (sta SampleTictailAuth) GetAccessToken() string {
	return sta.accessToken
}
