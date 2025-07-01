package proxy

import (
	"fmt"
	"sync"
	"time"
)

// Subject interface
type Image interface {
	Display()
}

// RealSubject – expensive to create
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	fmt.Printf("Loading image from disk: %s\n", filename)
	time.Sleep(1 * time.Second) // simulate expensive load
	return &RealImage{filename}
}

func (ri *RealImage) Display() {
	fmt.Printf("Displaying image: %s\n", ri.filename)
}

// Proxy – defers RealImage creation until Display is called
type ProxyImage struct {
	filename  string
	realImage *RealImage
	loadOnce  sync.Once
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (pi *ProxyImage) Display() {
	// load image only once
	pi.loadOnce.Do(func() {
		pi.realImage = NewRealImage(pi.filename)
	})
	pi.realImage.Display()
}

// Client
func RunImageLoaderDemo() {
	img := NewProxyImage("photo.png")
	// image not loaded yet
	fmt.Println("Proxy created; no load yet")
	img.Display() // loads and displays
	img.Display() // only displays
}
