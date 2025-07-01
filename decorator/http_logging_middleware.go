package decorator

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Component interface
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// Concrete Component
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Decorator
type LoggingMiddleware struct {
	next Handler
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lm.next.ServeHTTP(w, r) // delegate to wrapped handler
	log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
}

// Helper to chain decorators
func WithLogging(h Handler) Handler {
	return &LoggingMiddleware{next: h}
}

func main() {
	// Build the chain:
	// HelloHandler → LoggingMiddleware
	var handler Handler = &HelloHandler{}
	handler = WithLogging(handler)

	// Adapt to http.Handler
	http.Handle("/", handler)
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
