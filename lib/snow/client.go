package snow

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
)

// Config ...
type Config struct {
	URL *url.URL
}

// Client ...
type Client interface {
	Connect(chg Change) ([]string, []string)
}

type client struct{}

// NewSNOWClient ...
func NewSNOWClient(sc Config) (Client, error) {
	if sc.URL == nil {
		return nil, errors.New("error")
	}
	return &client{}, nil
}

// Connect ...
func (c *client) Connect(chg Change) ([]string, []string) {
	var payload []byte
	// if request.hasProxy {
	// 	payload = []byte(`{"target":"` + request.host + `:` + request.port + `", "http_proxy":"` + request.proxyHost + `:` + request.proxyPort + `"}`)
	// } else {
	// 	payload = []byte(`{"target":"` + request.host + `:` + request.port + `"}`)
	// }
	payload = []byte(`{"I'm a Change!"}`)
	req, err := http.NewRequest("POST", "http://abc.com", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, []string{"Unable to hit url", err.Error()}
	}
	defer resp.Body.Close()

	//decoder := json.NewDecoder(resp.Body)
	// decoder := json.NewDecoder(resp.Body)
	// var body wicResponse
	// decodeErr := decoder.Decode(&body)
	// if decodeErr != nil {
	// 	return nil, []string{"Invalid response from willitconnect: ", decodeErr.Error()}
	// }
	// var response []string
	// if body.CanConnect {
	// 	response = []string{"I am able to connect"}
	// } else {
	// 	response = []string{"I am unable to connect"}
	// }
	return []string{""}, nil
}
