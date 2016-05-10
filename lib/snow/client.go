package snow

import (
	"errors"
	"net/url"
)

type SNOWConfig struct {
	URL *url.URL
}

type Client interface {
}

type client struct{}

func NewSNOWClient(sc SNOWConfig) (Client, error) {
	if sc.URL == nil {
		return nil, errors.New("error")
	}

	return client{}, nil

}
