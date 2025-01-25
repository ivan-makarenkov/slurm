package main

import (
	"context"
	"fmt"
	"net/url"
)

const (
	health = "/health"
)

const (
	noData       = "No data"
	dataTemplate = "Overall status is %s, with service_id %s mysql component is %s"
)

func main() {
	// TODO: код писать здесь
}

type Client struct {
	u string
}

// NewClient constructor for healthchecker,
// parse url and use baseURL.
func NewClient(uri string) (*Client, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("parse uri: %w", err)
	}

	baseURL := u.Scheme + "://" + u.Host

	if len(uri) == 0 {
		baseURL = ""
	}

	return &Client{u: baseURL}, nil
}

type HealthCheck struct {
	// TODO: код писать здесь
}

func (c *Client) getHealth(ctx context.Context) string {
	// TODO: код писать здесь
	return "todo: replace me"
}
