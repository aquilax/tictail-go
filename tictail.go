package tictail

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	API_URL = "https://api.tictail.com"
	API_VER = "v1"

	API_GET = "GET"

	API_STORES = "stores"
)

type TictailAuth interface {
	GetAccessToken() string
}

type TictailLogger interface {
	Panic(v ...interface{})
	Print(v ...interface{})
}

type Tictail struct {
	ta     TictailAuth
	client *http.Client
	log    TictailLogger
}

func NewTictail(ta TictailAuth, logger TictailLogger) *Tictail {
	return &Tictail{
		ta,
		&http.Client{},
		logger,
	}
}

func (t Tictail) prepareRequestHeaders(req *http.Request) {
	// Add auth header
	// ref: https://tictail.com/developers/documentation/authentication/#Make_your_first_API_call
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.ta.GetAccessToken()))
	// API expects content-type to be application/json
	// ref: https://tictail.com/developers/documentation/api-reference/#always-set-the-required-headers
	req.Header.Add("Content-Type", "application/json")
}

func (t Tictail) makeRequest(method string, url string, data io.Reader) (int, []byte) {
	req, err := http.NewRequest(method, url, data)
	t.logError(err)
	t.prepareRequestHeaders(req)
	resp, err := t.client.Do(req)
	t.logError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.logError(err)
	return resp.StatusCode, body
}

func (t Tictail) createUrl(resource, params string) string {
	return API_URL + "/" + API_VER + "/" + resource + "/" + params
}

func (t Tictail) GetStore(id string) (TStore, TError) {
	var terror TError
	var tstore TStore
	url := t.createUrl(API_STORES, id)
	code, body := t.makeRequest(API_GET, url, nil)
	if code < 200 || code > 299 {
		t.logError(json.Unmarshal(body, &terror))
	} else {
		t.logError(json.Unmarshal(body, &tstore))
	}
	return tstore, terror
}

func (t Tictail) logError(err error) {
	if err != nil {
		t.log.Print(err)
	}
}
