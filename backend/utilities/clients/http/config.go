package http

import (
	"io"
	"net/http"
)

type Client interface {
	Do(method Method, url string, body io.Reader, headers map[string]string) (*http.Response, error)
}

type ClientImpl struct {
	httpClient    *http.Client
	globalHeaders map[string]string
}

func NewHttpClient(globalHeaders map[string]string) Client {
	return &ClientImpl{httpClient: http.DefaultClient, globalHeaders: globalHeaders}
}

func (c *ClientImpl) setHeaders(request *http.Request, headers map[string]string) {
	for k, v := range headers {
		request.Header.Set(k, v)
	}
}

func (c *ClientImpl) Do(method Method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(string(method), url, body)
	if err != nil {
		return nil, err
	}

	c.setHeaders(req, defaultHeaders)
	c.setHeaders(req, c.globalHeaders)
	c.setHeaders(req, headers)

	return c.httpClient.Do(req)
}
