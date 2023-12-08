package requests

import (
	"bytes"
	"fmt"
	"github.com/galagoshin-com/GoLogger/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Method string

const (
	GET    = Method("GET")
	POST   = Method("POST")
	DELETE = Method("DELETE")
	PUT    = Method("PUT")
	PATCH  = Method("PATCH")
)

type Request struct {
	Method  Method
	Url     URL
	Data    url.Values
	Headers http.Header
	Cookies []*http.Cookie
	Timeout time.Duration
}

func (request *Request) Send() (*Response, error) {
	req, _ := http.NewRequest(string(request.Method), string(request.Url), bytes.NewBuffer([]byte(request.Data.Encode())))
	var client http.Client
	if request.Timeout == time.Duration(0) {
		client = http.Client{}
	} else {
		client = http.Client{
			Timeout: request.Timeout,
		}
	}
	req.Header = request.Headers
	for _, cookie := range request.Cookies {
		req.AddCookie(cookie)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	response := &Response{
		Body:       body,
		Header:     res.Header,
		Cookies:    res.Cookies(),
		Status:     res.Status,
		StatusCode: res.StatusCode,
	}
	logger.Debug(4, false, fmt.Sprintf("Request method: %s", request.Method))
	logger.Debug(4, false, fmt.Sprintf("Request URL: %s", request.Url))
	logger.Debug(5, false, fmt.Sprintf("Request data: %s", request.Data))
	logger.Debug(6, false, fmt.Sprintf("Request headers: %s", request.Headers))
	logger.Debug(4, false, fmt.Sprintf("Response status: %s", response.StatusCode))
	logger.Debug(4, false, fmt.Sprintf("Response status: %s", response.Body))
	logger.Debug(5, false, fmt.Sprintf("Response cookies: %s", response.Cookies))
	logger.Debug(6, false, fmt.Sprintf("Response headers: %s", response.Header))
	return response, nil
}
