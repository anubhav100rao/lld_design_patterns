package chain_of_responsibility

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HandlerFunc defines the signature for our chain handlers.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) bool

// Chain holds the sequence of handlers.
type Chain struct {
	handlers []HandlerFunc
	final    http.Handler
}

// NewChain initializes a new Chain.
func NewChain(final http.Handler, handlers ...HandlerFunc) *Chain {
	return &Chain{handlers: handlers, final: final}
}

// ServeHTTP executes handlers in order until one returns false, else calls final.
func (c *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range c.handlers {
		if proceed := handler(w, r); !proceed {
			return // stop chain
		}
	}
	c.final.ServeHTTP(w, r)
}

// Middleware examples

// LoggingMiddleware logs requests and always proceeds.
func LoggingMiddleware(w http.ResponseWriter, r *http.Request) bool {
	start := time.Now()
	log.Printf("Started %s %s", r.Method, r.URL.Path)
	// we could store start time in context for later
	fmt.Fprintf(w, "") // ensure header is written
	log.Printf("Completed in %v", time.Since(start))
	return true
}

// AuthMiddleware checks for a header and stops chain if unauthorized.
func AuthMiddleware(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("X-Auth-Token") != "secret" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}
	return true
}

// FinalHandler is the end of the chain.
var FinalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, you are authorized!")
})

func RunHTTPMiddlewareChain() {
	chain := NewChain(FinalHandler,
		LoggingMiddleware,
		AuthMiddleware,
	)
	http.Handle("/", chain)
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
