package proxy

import (
	"errors"
	"fmt"
)

// Subject interface
type Service interface {
	PerformAction(user string) error
}

// RealSubject
type RealService struct{}

func (rs *RealService) PerformAction(user string) error {
	fmt.Printf("Action performed for user %s\n", user)
	return nil
}

// Proxy – checks permissions before forwarding
type AuthProxy struct {
	service Service
	allowed map[string]bool
}

func NewAuthProxy(s Service, allowedUsers []string) *AuthProxy {
	perm := make(map[string]bool)
	for _, u := range allowedUsers {
		perm[u] = true
	}
	return &AuthProxy{service: s, allowed: perm}
}

func (p *AuthProxy) PerformAction(user string) error {
	if !p.allowed[user] {
		return errors.New("access denied for user: " + user)
	}
	// forward to real service
	return p.service.PerformAction(user)
}

// Client
func RunAuthProxyDemo() {
	real := &RealService{}
	proxy := NewAuthProxy(real, []string{"alice", "bob"})

	for _, user := range []string{"alice", "eve"} {
		if err := proxy.PerformAction(user); err != nil {
			fmt.Println(err)
		}
	}
}
