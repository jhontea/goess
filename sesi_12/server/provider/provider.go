package provider

import "pencairan_user/server/client"

type checkInterface interface {
	CheckUrl(string, chan string)
}

type checker struct{}

var Checker checkInterface = &checker{}

func (check *checker) CheckUrl(url string, c chan string) {
	resp, err := client.ClientCall.GetValue(url)
	if err != nil {
		c <- "cant_access_resource"
		return
	}

	if resp.StatusCode > 299 {
		c <- "no_match"
	}

	if resp.StatusCode == 200 {
		c <- url
	}
}
