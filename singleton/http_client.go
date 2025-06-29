package singleton

import (
	"net"
	"net/http"
	"time"
)

var (
	client *http.Client
)

// GetClient returns a singleton HTTP client with sane defaults.
func GetClient() *http.Client {
	once.Do(func() {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 5 * time.Second,
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}
	})
	return client
}
