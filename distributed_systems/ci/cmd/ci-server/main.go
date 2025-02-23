package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/mmiller-hairston/exploration/distributed_systems/ci/middleware"
	"github.com/mmiller-hairston/exploration/distributed_systems/ci/routes"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8081"`
}

func main() {
	cli := humacli.New(func(h humacli.Hooks, o *Options) {
		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))

		// API Mux
		a := http.NewServeMux()
		conf := huma.DefaultConfig("CI Server", "0.0.1")
		conf.Servers = []*huma.Server{
			{URL: "http://localhost:8081/api/v1"},
		}
		api := humago.New(a, conf)
		routes.Register(api, l)

		// Root Mux
		r := http.NewServeMux()
		r.Handle("/api/v1/", http.StripPrefix("/api/v1", a))

		s := http.Server{
			Addr:    fmt.Sprintf(":%d", o.Port),
			Handler: middleware.Logger(l)(r),
		}

		h.OnStart(func() {
			l.Info("Running...", slog.Int("port", o.Port))
			s.ListenAndServe()
		})

		h.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			s.Shutdown(ctx)
		})
	})

	cli.Run()
}
