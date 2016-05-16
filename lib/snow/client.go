package snow

import (
	"errors"
	"net/url"
)

// Config ...
type Config struct {
	URL *url.URL
}

// Client ...
type Client interface {
}

type client struct{}

// NewSNOWClient ...
func NewSNOWClient(sc Config) (Client, error) {
	if sc.URL == nil {
		return nil, errors.New("error")
	}

	return client{}, nil

}
