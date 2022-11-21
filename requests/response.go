package requests

import "net/http"

type Response struct {
	Body       []byte
	Header     http.Header
	Cookies    []*http.Cookie
	Status     string
	StatusCode int
}

func (response *Response) Text() string {
	return string(response.Body)
}
