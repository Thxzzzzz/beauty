package beauty

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rushteam/beauty/pkg/service/webserver"
)

func WithWebServer(addr string, mux http.Handler) Option {
	return func(app *App) {
		app.services = append(app.services, webserver.New(addr, mux))
	}
}

type RouteOption func(r *chi.Mux)

func WithWebServerChi(addr string, routes []Route, opts ...RouteOption) Option {
	r := chi.NewRouter()
	for _, v := range opts {
		v(r)
	}
	for _, v := range routes {
		if v.Method == "" {
			v.Method = http.MethodGet
		}
		r.Method(v.Method, v.URI, v.Handler)
	}
	s := webserver.New(addr, r)
	return func(app *App) {
		app.services = append(app.services, s)
	}
}
func WithChiMiddleware(middlewares ...func(http.Handler) http.Handler) RouteOption {
	return func(r *chi.Mux) {
		for _, v := range middlewares {
			r.Use(v)
		}
	}
}
