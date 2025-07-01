package proxy

import (
	"fmt"
	"net/rpc"
)

// Subject interface
type KeyValueStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

// RealSubject (server‑side implementation)
type KVServer struct {
	store map[string]string
}

func (s *KVServer) Get(args string, reply *string) error {
	*reply = s.store[args]
	return nil
}

func (s *KVServer) Set(kv [2]string, reply *bool) error {
	s.store[kv[0]] = kv[1]
	*reply = true
	return nil
}

// Proxy (client‑side stub)
type KVClientProxy struct {
	client *rpc.Client
}

func NewKVClientProxy(address string) (*KVClientProxy, error) {
	c, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &KVClientProxy{client: c}, nil
}

func (p *KVClientProxy) Get(key string) (string, error) {
	var reply string
	if err := p.client.Call("KVServer.Get", key, &reply); err != nil {
		return "", err
	}
	return reply, nil
}

func (p *KVClientProxy) Set(key, value string) error {
	var ok bool
	if err := p.client.Call("KVServer.Set", [2]string{key, value}, &ok); err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to set %s", key)
	}
	return nil
}

// Client
func RunRemoteProxyDemo() {
	// assume RPC server is already running and registered as "KVServer"
	proxy, err := NewKVClientProxy("127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	proxy.Set("foo", "bar")
	val, _ := proxy.Get("foo")
	fmt.Println("foo =", val)
}
