package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

//Registry ..
type Registry interface {
	Register(ctx context.Context, s *Service, ttl time.Duration) error
	Deregister(ctx context.Context, s *Service) error
	Discover(ctx context.Context, naming string) ([]*Service, error)
	Services(ctx context.Context, naming string) ([]*Service, error)
}

//Optipns ..
type Optipns func(s *Service)

//WithNamespace ..
func WithNamespace(p string) Optipns {
	return func(s *Service) {
		s.Namespace = p
	}
}

//WithKind ..
func WithKind(p string) Optipns {
	return func(s *Service) {
		s.Kind = p
	}
}

//WithVersion ..
func WithVersion(p string) Optipns {
	return func(s *Service) {
		s.Version = p
	}
}

//WithName ..
func WithName(p string) Optipns {
	return func(s *Service) {
		s.Name = p
	}
}

//NewService ..
func NewService(opts ...Optipns) *Service {
	s := &Service{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

//Service info
type Service struct {
	Namespace string            `json:"namespace"`
	Kind      string            `json:"kind"`
	Name      string            `json:"name"`
	ID        string            `json:"id"`
	Version   string            `json:"var"`
	Address   string            `json:"addr"`
	Labels    map[string]string `json:"labels"`
	// *Node
	// Metadata map[string]string `json:"metadata"`
	// Endpoints []*Endpoint       `json:"endpoints"`
	// Nodes     []*Node           `json:"nodes"`
}

//Node info
type Node struct {
	ID      string            `json:"id"`
	Version string            `json:"ver"`
	Address string            `json:"addr"`
	Labels  map[string]string `json:"labels"`
}

//Unmarshal ..
func (n *Node) Unmarshal(v []byte) error {
	return json.Unmarshal(v, n)
}

//naming ..
func naming(namespace, kind, name string) string {
	return fmt.Sprintf("%v/%v/%v/%v/", prefix, namespace, kind, name)
}

//String ..
func (s *Service) String() string {
	return fmt.Sprintf("%v%v", naming(s.Namespace, s.Kind, s.Name), s.ID)
}

//Marshal ...
func (s *Service) Marshal() []byte {
	v, _ := json.Marshal(s)
	return v
}

//Unmarshal ..
func Unmarshal(v []byte) *Service {
	var s *Service
	json.Unmarshal(v, s)
	return s
}

//Endpoint ..
// type Endpoint struct {
// 	Name     string            `json:"name"`
// 	Request  *Value            `json:"request"`
// 	Response *Value            `json:"response"`
// 	Metadata map[string]string `json:"metadata"`
// }
