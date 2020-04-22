package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	var (
		resp    *http.Response
		err     error
		content []byte
	)
	if resp, err = http.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	if content, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return content, nil
}
