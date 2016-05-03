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

func NewSNOWClient(sc SNOWConfig) (Client, error) {
	return nil, errors.New("error")
}
