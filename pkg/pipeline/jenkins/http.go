package jenkins

import (
	"io/ioutil"
	"net/http"
)

type httpCall interface {
	httpGET(URL string) ([]byte, error)
}

type client struct {
	cli *http.Client
}

// httpGET GET http call
func (c *client) httpGET(URL string) ([]byte, error) {
	response, err := c.cli.Get(URL)
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
