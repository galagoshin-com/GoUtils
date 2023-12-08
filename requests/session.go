package requests

import (
	"fmt"
	"github.com/galagoshin-com/GoLogger/logger"
	"net/http"
)

type Session struct {
	Cookies []*http.Cookie
}

func (session *Session) Init() {
	session.Cookies = []*http.Cookie{}
}

func (session *Session) Request(request Request) (*Response, error) {
	response, err := request.Send()
	if err != nil {
		return nil, err
	}
	cookieIndex := func(cookie_to_find *http.Cookie) int {
		for id, cookie := range session.Cookies {
			if cookie == cookie_to_find {
				return id
			}
		}
		return -1
	}
	for _, cookie := range response.Cookies {
		if index := cookieIndex(cookie); index == -1 {
			session.Cookies = append(session.Cookies, cookie)
		} else {
			session.Cookies[index] = cookie
		}
	}
	logger.Debug(5, false, fmt.Sprintf("Session cookies: %s", session.Cookies))
	return response, nil
}
