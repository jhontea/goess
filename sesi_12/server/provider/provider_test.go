package provider

import (
	"errors"
	"net/http"
	"testing"

	"pencairan_user/server/client"

	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc func(url string) (*http.Response, error)
)

type clientMock struct{}

//mocking the client call:
func (cm *clientMock) GetValue(url string) (*http.Response, error) {
	return getRequestFunc(url)
}

//When the api call is successful and the desired result is gotten
func TestCheckUrls_Success(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}
	client.ClientCall = &clientMock{}
	url := "https://twitter.com/hacktiv8id"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	result := <-ch
	assert.NotNil(t, result)
	assert.EqualValues(t, "https://twitter.com/hacktiv8id", result)
}

//When the api call is not successful, maybe there is no internet connection
func TestCheckUrls_Not_Existent_Url(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return nil, errors.New("there is an error here")
	}
	client.ClientCall = &clientMock{}
	url := "https://invalid_url/hacktiv8id"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	err := <-ch
	assert.NotNil(t, err)
	assert.EqualValues(t, "cant_access_resource", err)
}

//When the api call is successful, but the desire result is not produced
func TestCheckUrls_Username_Dont_Exist(t *testing.T) {
	getRequestFunc = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
		}, nil
	}
	client.ClientCall = &clientMock{}
	url := "https://twitter.com/random_xxxxx_yyyy"
	ch := make(chan string)
	go Checker.CheckUrl(url, ch)
	result := <-ch
	assert.NotNil(t, result)
	assert.EqualValues(t, "no_match", result)
}
