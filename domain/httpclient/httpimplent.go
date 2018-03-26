package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Implement struct {
}

func (i Implement) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http response failed")
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil ReadAll failed")
	}
	return response, nil
}

func (i Implement) Post(url string) ([]byte, error) {
	return nil, nil
}
