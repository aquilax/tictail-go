package tictail

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type TictailAuth interface {
	getAuthCode() string
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.ta.getAuthCode()))
	// API expects content-type to be application/json
	// ref: https://tictail.com/developers/documentation/api-reference/#always-set-the-required-headers
	req.Header.Add("Content-Type", "application/json")
}

func (t Tictail) makeRequest(method string, url string, data io.Reader) (int, []byte) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		t.log.Print(err)
	}
	t.prepareRequestHeaders(req)
	resp, err := t.client.Do(req)
	if err != nil {
		t.log.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.log.Print(err)
	}
	return resp.StatusCode, body
}
