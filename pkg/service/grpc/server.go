package grpc

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rushteam/beauty/pkg/config"
	"github.com/rushteam/beauty/pkg/log"
	"github.com/rushteam/beauty/pkg/registry"
)

//ServiceKind ..
const ServiceKind = "grpc"

//Build create a rest service with the name
func Build(name string) (*Server, error) {
	s := &Server{
		service: &registry.Service{
			Kind: ServiceKind,
			Name: name,
		},
		Mode: gin.DebugMode,
		Addr: ":grpc",
	}
	if conf, err := config.New(config.Env(), name); err == nil {
		s.Mode = conf.GetString(ServiceKind + ".mode")
		s.Addr = conf.GetString(ServiceKind + ".addr")
	} else {
		log.Warn("no config file...")
	}
	if len(s.Mode) > 0 {
		gin.SetMode(s.Mode)
	}

	s.Engine = gin.New()
	s.Server = &http.Server{
		Handler: s.Engine,
	}
	return s, nil
}

//Server ..
type Server struct {
	*gin.Engine
	Server  *http.Server
	listen  *net.Listener
	service *registry.Service
	Mode    string
	Addr    string
}

//Start ..
func (s *Server) Start(ctx context.Context) error {
	// log.Logger.Info("start with", ServiceKind, s)
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	if err := s.Server.Serve(ln); err != http.ErrServerClosed {
		return err
	}
	return nil
}

//Stop ..
func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

//Service ..
func (s *Server) Service() *registry.Service {
	return s.service
}