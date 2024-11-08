package main

import (
	"context"
	"fmt"
	"go-dependency-injection/pkg/handlers"
	"go-dependency-injection/pkg/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			handlers.NewEchoHandler,
			NewServeMux,
			zap.NewExample,
			services.NewPyramidService,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(echo *handlers.EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
