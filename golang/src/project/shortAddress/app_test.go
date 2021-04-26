package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

// -----------  测试用例

const (
	expTime       = 60
	longURL       = "https://www.google.com"
	shortLink     = "1"
	shortLinkInfo = `{"url":"https://www.google.com","created_at":"2021-04-26 15:21:53.545233 +0800 CST m=+2.520314494","expiration_in_minutes":60}`
)

type storageMock struct {
	mock.Mock
}

var app App
var mockR *storageMock

func (s *storageMock) Shorten(url string, exp int64) (string, error) {
	args := s.Called(url, exp)
	return args.String(0), args.Error(1)
}

func (s *storageMock) Unshorten(eid string) (string, error) {
	args := s.Called(eid)
	return args.String(0), args.Error(1)
}

func (s *storageMock) ShortlinkInfo(eid string) (interface{}, error) {
	args := s.Called(eid)
	return args.String(0), args.Error(1)
}

func init() {
	app = App{}
	mockR = new(storageMock)
	app.Initialize(&Env{S: mockR})
}

func TestCreateShortlink(t *testing.T) {
	var jsonStr = []byte(`{
		"url": "https://www.google.com",
		"expiration_in_minutes": 60}`)
	req, err := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal("Shoule be able to create a request.", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 预置预期的结果
	mockR.On("Shorten", longURL, int64(expTime)).Return(shortLink, nil).Once()
	rw := httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusCreated {
		t.Fatalf("Excepted receive %d. Got %d", http.StatusCreated, rw.Code)
	}
	resp := struct {
		Shortlink string `json:"shortlink"`
	}{}
	if err := json.NewDecoder(rw.Body).Decode(&resp); err != nil {
		t.Fatalf("Shoule decode the response")
	}
}

func TestRedircet(t *testing.T) {
	r := fmt.Sprintf("/%s", shortLink)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Fatal("Should be albe to create a request.", err)
	}

	mockR.On("Unshorten", shortLink).Return(longURL, nil).Once()
	rw := httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)
	if rw.Code != http.StatusTemporaryRedirect {
		t.Fatalf("Excepted receive %d. Got %d", http.StatusTemporaryRedirect, rw.Code)
	}
}
