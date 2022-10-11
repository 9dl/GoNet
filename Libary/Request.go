package GoNet

import (
	"io"
	"net/http"
)

func AddHeader(w *http.Request, key string, value string) {
	w.Header.Add(key, value)
}

func SetUserAgent(w *http.Request, agent string) {
	w.Header.Add("User-Agent", agent)
}

func GetResponse(res *http.Response) string {
	Response, _ := io.ReadAll(res.Body)
	return string(Response)
}

func GetRequest(URL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	return req, err
}

func PostRequest(URL string, ContentType string, Body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", URL, Body)
	req.Header.Add("Content-Type", ContentType)
	return req, err
}

func NewClient(AutoRedirect bool) *http.Client {
	client := &http.Client{}
	if AutoRedirect == true {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	return client
}

func NewResponse(Client *http.Client, Request *http.Request) (*http.Response, error) {
	return Client.Do(Request)
}
