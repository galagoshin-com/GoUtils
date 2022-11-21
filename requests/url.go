package requests

import (
	"net/url"
)

type URL string

func (addr URL) Parse() (url.Values, error) {
	u, err := url.Parse(string(addr))
	if err != nil {
		return nil, err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	return q, nil
}
