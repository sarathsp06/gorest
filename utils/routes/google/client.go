package google

import (
	"googlemaps.github.io/maps"
)

// Client represents a google client
type Client struct {
	*maps.Client
	Key string
}

// New returns a new client
func New(key string) (*Client, error) {
	client, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return nil, err
	}
	return &Client{Client: client, Key: key}, nil
}
