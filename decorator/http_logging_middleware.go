package decorator

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

type LoggingMiddleware struct {
	next Handler
}

func (lm *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lm.next.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL, time.Since(start))
}

func WithLogging(h Handler) Handler {
	return &LoggingMiddleware{next: h}
}

func RunHttpLoggingMiddlewareDemo() {
	var handler Handler = &HelloHandler{}
	handler = WithLogging(handler)

	http.Handle("/", handler)
	log.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
