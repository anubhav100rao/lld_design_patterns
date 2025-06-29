package builder

import (
	"fmt"
	"log"
	"net/url"
	"time"
)

type HTTPRequest struct {
	Method      string
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
	Retries     int
	Backoff     time.Duration
	Timeout     time.Duration
}

type RequestBuilder interface {
	SetMethod(method string) RequestBuilder
	SetURL(url string) RequestBuilder
	AddHeader(key, val string) RequestBuilder
	AddQueryParam(key, val string) RequestBuilder
	SetBody(body []byte) RequestBuilder
	SetRetries(count int, backoff time.Duration) RequestBuilder
	SetTimeout(timeout time.Duration) RequestBuilder
	Build() (*HTTPRequest, error)
}

type httpRequestBuilder struct {
	req HTTPRequest
	err error
}

func NewRequestBuilder() RequestBuilder {
	return &httpRequestBuilder{
		req: HTTPRequest{
			Headers:     make(map[string]string),
			QueryParams: make(map[string]string),
		},
	}
}

func (b *httpRequestBuilder) SetMethod(m string) RequestBuilder {
	if b.err != nil {
		return b
	}
	if m == "" {
		b.err = fmt.Errorf("method cannot be empty")
	}
	b.req.Method = m
	return b
}

func (b *httpRequestBuilder) SetURL(u string) RequestBuilder {
	if b.err != nil {
		return b
	}
	_, parseErr := url.ParseRequestURI(u)
	if parseErr != nil {
		b.err = parseErr
		return b
	}
	b.req.URL = u
	return b
}

func (b *httpRequestBuilder) AddHeader(k, v string) RequestBuilder {
	if b.err == nil {
		b.req.Headers[k] = v
	}
	return b
}

func (b *httpRequestBuilder) AddQueryParam(k, v string) RequestBuilder {
	if b.err == nil {
		b.req.QueryParams[k] = v
	}
	return b
}

func (b *httpRequestBuilder) SetBody(body []byte) RequestBuilder {
	if b.err == nil {
		b.req.Body = body
	}
	return b
}

func (b *httpRequestBuilder) SetRetries(count int, backoff time.Duration) RequestBuilder {
	if b.err == nil {
		if count < 0 {
			b.err = fmt.Errorf("retries cannot be negative")
		} else {
			b.req.Retries = count
			b.req.Backoff = backoff
		}
	}
	return b
}

func (b *httpRequestBuilder) SetTimeout(t time.Duration) RequestBuilder {
	if b.err == nil {
		if t <= 0 {
			b.err = fmt.Errorf("timeout must be > 0")
		} else {
			b.req.Timeout = t
		}
	}
	return b
}

func (b *httpRequestBuilder) Build() (*HTTPRequest, error) {
	if b.err != nil {
		return nil, b.err
	}
	return &b.req, nil
}

func RunHTTPBuilderDemo() {
	req, err := NewRequestBuilder().
		SetMethod("POST").
		SetURL("https://api.example.com/v1/orders").
		AddHeader("Content-Type", "application/json").
		AddQueryParam("dry_run", "true").
		SetBody([]byte(`{"item":"book","qty":1}`)).
		SetRetries(3, 500*time.Millisecond).
		SetTimeout(5 * time.Second).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	// execute req...
	fmt.Printf("Built HTTP Request: %+v\n", req)
}
