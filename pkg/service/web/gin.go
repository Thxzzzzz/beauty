package web

import (
	"context"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rushteam/beauty/pkg/config"
	"github.com/rushteam/beauty/pkg/log"
)

//ServiceKind ..
const ServiceKind = "web.gin"

//New new a rest service with the name
func New(name string) (*Rest, error) {
	s := &Rest{
		Mode: gin.DebugMode,
		Addr: ":http",
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

//Rest ..
type Rest struct {
	*gin.Engine
	Server *http.Server
	Mode   string
	Addr   string
}

//Start ..
func (s *Rest) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	if err := s.Server.Serve(ln); err != http.ErrServerClosed {
		return err
	}
	return err
}

//Stop ..
func (s *Rest) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}

// String ..
func (s *Rest) String() string {
	return "web"
}