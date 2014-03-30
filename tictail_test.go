package tictail

import (
	"testing"
)

type AuthMock struct{}

func (am AuthMock) GetAccessToken() string {
	return "access_token"
}

type LoggerMock struct{}

func (l LoggerMock) Panic(v ...interface{}) {
	//
}
func (l LoggerMock) Print(v ...interface{}) {
	print(v)
}

func TestGetStore(t *testing.T) {
	am := AuthMock{}
	lm := LoggerMock{}
	ti := NewTictail(am, lm)
	store, err := ti.GetStore("store_1")
	print(store.Id)
	print(err.Message)
}
