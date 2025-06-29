package prototype

import (
	"bytes"
	"net/http"
	"time"
)

// Prototype: wraps http.Request and common settings
type RequestPrototype struct {
	Method  string
	URL     string
	Headers http.Header
	Timeout time.Duration
	Body    []byte
}

// Clone performs a deep copy
func (p *RequestPrototype) Clone() *RequestPrototype {
	hdr := make(http.Header, len(p.Headers))
	for k, v := range p.Headers {
		hdr[k] = append([]string{}, v...)
	}
	bodyCopy := append([]byte{}, p.Body...)
	return &RequestPrototype{
		Method:  p.Method,
		URL:     p.URL,
		Headers: hdr,
		Timeout: p.Timeout,
		Body:    bodyCopy,
	}
}

// Build converts prototype into an *http.Request
func (p *RequestPrototype) Build() (*http.Request, *http.Client, error) {
	req, err := http.NewRequest(p.Method, p.URL, bytes.NewReader(p.Body))
	if err != nil {
		return nil, nil, err
	}
	req.Header = p.Headers.Clone()
	client := &http.Client{Timeout: p.Timeout}
	return req, client, nil
}

// Usage
func ExampleAPI() error {
	base := &RequestPrototype{
		Method:  "POST",
		URL:     "https://api.example.com/v2/items",
		Headers: http.Header{"Authorization": {"Bearer <token>"}, "Content-Type": {"application/json"}},
		Timeout: 5 * time.Second,
	}

	// Clone and customize for each payload
	for _, payload := range []string{`{"id":1}`, `{"id":2}`} {
		reqProto := base.Clone()
		reqProto.Body = []byte(payload)

		req, client, err := reqProto.Build()
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		resp.Body.Close()
	}
	return nil
}
