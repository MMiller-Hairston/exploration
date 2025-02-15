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
	"github.com/mmiller-hairston/exploration/distributed_systems/ci/routes"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8080"`
}

func main() {
	cli := humacli.New(func(h humacli.Hooks, o *Options) {
		l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		r := http.NewServeMux()
		api := humago.New(r, huma.DefaultConfig("CI Server", "0.0.1"))

		routes.Register(api, l)

		s := http.Server{
			Addr:    fmt.Sprintf(":%d", o.Port),
			Handler: r,
		}

		h.OnStart(func() {
			fmt.Printf("Running...")
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
