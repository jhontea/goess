package client

import "net/http"

var (
	ClientCall HTTPClient = &clientCall{}
)

type HTTPClient interface {
	GetValue(url string) (*http.Response, error)
}

type clientCall struct {
	Client http.Client
}

func (cli *clientCall) GetValue(url string) (*http.Response, error) {
	res, err := cli.Client.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
